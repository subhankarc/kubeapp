package main

import (
	"log"
	"net/http"

	"github.com/smjn/ipl18/backend"
	"github.com/smjn/ipl18/backend/config"
)

func main() {
	log.Println("Starting server on port", config.GetConfig().AppPort)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+config.GetConfig().AppPort, backend.SetupAndGetRouter()))
}
