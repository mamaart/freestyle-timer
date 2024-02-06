package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mamaart/freestyle-timer/internal/api"
	"github.com/mamaart/freestyle-timer/internal/app"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	a := api.New(app.New())

	r.HandleFunc("/new", a.NewSession)
	r.HandleFunc("/destroy", a.DestroySession)
	r.HandleFunc("/start/{id}", a.StartTimer)
	r.HandleFunc("/pause/{id}", a.PauseTimer)
	r.HandleFunc("/state", a.GetState)

	fmt.Println("Serving slack timer on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
