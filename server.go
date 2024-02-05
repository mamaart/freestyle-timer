package main

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "missing session name", http.StatusNotAcceptable)
			return
		}
	})

	fmt.Println("Serving slack timer on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
