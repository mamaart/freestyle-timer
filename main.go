package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mamaart/freestyle-timer/internal/session"
)

func main() {
	s := session.New(time.Minute)

	if err := s.Start(1); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s.String())
	}

	if err := s.Pause(1); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s.String())
	}

	if err := s.Start(2); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(s.String())
	}
}
