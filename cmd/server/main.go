package main

import (
	"fmt"
	"log"

	"github.com/mamaart/freestyle-timer/internal/api"
	"github.com/mamaart/freestyle-timer/internal/app"
)

func main() {
	a := api.New(app.New())
	fmt.Println("Serving slack timer on port :8080")
	log.Fatal(a.Serve())
}
