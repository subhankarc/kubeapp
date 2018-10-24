package config

import (
	"log"
	"os"

	"github.com/smjn/ipl18/backend/models"
)

var serverConfig models.Config

var GetConfig = func() models.Config {
	return serverConfig
}

var GetDBConfig = func() models.DBConfig {
	return serverConfig.DB
}

func init() {
	log.Println("Parsing config from env...")

	if os.Getenv("POSTGRES_DB") != "" {
		serverConfig.DB.DBName = os.Getenv("POSTGRES_DB")
	} else {
		log.Println("error parsing config")
		os.Exit(1)
	}
	if os.Getenv("POSTGRES_USER") != "" {
		serverConfig.DB.DBUser = os.Getenv("POSTGRES_USER")
	} else {
		log.Println("error parsing config")
		os.Exit(1)
	}
	if os.Getenv("POSTGRES_PASSWORD") != "" {
		serverConfig.DB.DBPassword = os.Getenv("POSTGRES_PASSWORD")
	} else {
		log.Println("error parsing config")
		os.Exit(1)
	}
	if os.Getenv("POSTGRES_HOST") != "" {
		serverConfig.DB.Host = os.Getenv("POSTGRES_HOST")
	} else {
		log.Println("error parsing config")
		os.Exit(1)
	}
	serverConfig.DB.Port = "5432"
	log.Println(serverConfig)
}
