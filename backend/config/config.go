package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"

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
	if os.Getenv("app_config") != "" {
		if err := json.NewDecoder(strings.NewReader(os.Getenv("app_config"))).Decode(&serverConfig); err != nil {
			log.Println("error parsing config", err.Error())
			os.Exit(1)
		}
	} else {
		log.Println("Config: app_config not found in env, exiting")
		os.Exit(1)
	}
	log.Println(serverConfig)
}
