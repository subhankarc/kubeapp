package main

import (
	"log"
	"net/http"

	"github.com/smjn/ipl18/backend"
)

func main() {
	log.Println("Starting server on port", 5000)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", backend.SetupAndGetRouter()))
}
