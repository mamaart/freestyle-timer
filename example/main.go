package main

import (
	"log"

	"github.com/mamaart/freestyle-timer/internal/session"
)

func main() {
	if err := session.TEST(); err != nil {
		log.Fatal(err)
	}
}
