package main

import (
	"log"
	"net/http"

	"github.com/smjn/kubeapp/backend"
)

func main() {
	log.Println("Starting server on port", 8080)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", backend.SetupAndGetRouter()))
}
