package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mamaart/freestyle-timer/internal/app"
	"github.com/mamaart/freestyle-timer/models"
)

type Api struct {
	app *app.App
}

func New(app *app.App) *Api {
	return &Api{app: app}
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
