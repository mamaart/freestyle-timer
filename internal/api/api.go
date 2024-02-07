package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mamaart/freestyle-timer/internal/app"
	"github.com/mamaart/freestyle-timer/models"
)

type Api struct {
	app      *app.App
	upgrader *websocket.Upgrader
}

func New(app *app.App) *Api {
	return &Api{
		app:      app,
		upgrader: &websocket.Upgrader{},
	}
}

func (a *Api) Serve() error {
	r := mux.NewRouter()
	r.HandleFunc("/new", a.NewSession)
	r.HandleFunc("/destroy", a.DestroySession)
	r.HandleFunc("/start/{id}", a.StartTimer)
	r.HandleFunc("/pause/{id}", a.PauseTimer)
	r.HandleFunc("/state", a.GetState)
	r.HandleFunc("/listen/{id}", a.Listen)
	return http.ListenAndServe(":8080", r)
}

func (a *Api) NewSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var opt models.Opt

	if err := json.NewDecoder(r.Body).Decode(&opt); err != nil {
		http.Error(w, "failed decoding session options", http.StatusNotAcceptable)
		return
	}

	fmt.Println(opt)

	if err := a.app.NewSession(opt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) DestroySession(w http.ResponseWriter, r *http.Request) {
	if err := a.app.DestroyCurrentSession(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) StartTimer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "invalid athlete id", http.StatusNotAcceptable)
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid athlete id", http.StatusNotAcceptable)
		return
	}

	if err := a.app.StartTimer(i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) PauseTimer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "invalid athlete id", http.StatusNotAcceptable)
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid athlete id", http.StatusNotAcceptable)
		return
	}

	if err := a.app.PauseTimer(i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) GetState(w http.ResponseWriter, r *http.Request) {
	state, err := a.app.GetState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(state))
}

func (a *Api) Listen(w http.ResponseWriter, r *http.Request) {
	conn, err := a.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	id := mux.Vars(r)["id"]
	if id == "" {
		conn.WriteMessage(websocket.TextMessage, []byte("invalid athlete id"))
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("invalid athlete id"))
		return
	}
	ch, err := a.app.Listen(i)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	for t := range ch {
		x := fmt.Sprintf("%02d:%02d", int(t.Minutes()), int(t.Seconds())%60)
		conn.WriteMessage(websocket.TextMessage, []byte(x))
	}
}
