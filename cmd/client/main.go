package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/mamaart/freestyle-timer/models"
	"github.com/mamaart/freestyle-timer/pkg/image"
)

func main() {
	// newSession()
	go listenPlayer("martin", 1)
	go listenPlayer("john", 2)
	<-make(chan bool)
}

func listenPlayer(name string, i int) {
	url := fmt.Sprintf("ws://localhost:8080/listen/%d", i)
	conn, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)

	ch := make(chan int)

	s := image.NewImageState(name, ch)
	if i == 1 {
		go s.Run("one")
	} else {
		go s.Run("two")
	}
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		x, err := strconv.Atoi(string(data))
		if err != nil {
			log.Fatal(err)
		}
		ch <- x
	}
}

func newSession() {
	opt := models.Opt{
		SessionTitle:   "scandinavian finals",
		AthleteOneName: "martin",
		AthleteTwoName: "john",
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
}
