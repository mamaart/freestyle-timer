package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/mamaart/freestyle-timer/models"
)

func main() {
	opt := models.Opt{
		SessionTitle:   "first session",
		AthleteOneName: "martin",
		AthleteTwoName: "bob",
	}

	buf := bytes.Buffer{}

	json.NewEncoder(&buf).Encode(&opt)

	resp, err := http.DefaultClient.Post("http://localhost:8080/new", "application/json", &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	if resp.StatusCode != 200 {
		log.Fatal(errors.New("not success"))
	}
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/"}
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
