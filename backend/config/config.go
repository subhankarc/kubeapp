package config

import (
	"log"
        "github.com/cloudfoundry-community/go-cfenv"
	"github.com/smjn/kubeapp/backend/models"
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

        appEnv, _ := cfenv.Current()
        pgService, _ := appEnv.Services.WithName("ali-pg")
	log.Println("PostgreSQL service: ", pgService)

	dbname, _ := pgService.CredentialString("dbname")
	if dbname != "" {
		serverConfig.DB.DBName = dbname
	} else {
		log.Println("error parsing config POSTGRES_DB")
		// os.Exit(1)
	}

	username, _ := pgService.CredentialString("username")
	if username != "" {
		serverConfig.DB.DBUser = username
	} else {
		log.Println("error parsing config POSTGRES_USER")
		// os.Exit(1)
	}

	password, _ := pgService.CredentialString("password")
	if password != "" {
		serverConfig.DB.DBPassword = password
	} else {
		log.Println("error parsing config POSTGRES_PASSWORD")
		// os.Exit(1)
	}

	hostname, _ := pgService.CredentialString("hostname")
	if hostname != "" {
		serverConfig.DB.Host = hostname
	} else {
		log.Println("error parsing config POSTGRES_SERVICE_HOST")
		// os.Exit(1)
	}

        port, _ := pgService.CredentialString("port")
        if port != "" {
                serverConfig.DB.Port = port
        } else {
	        serverConfig.DB.Port = "5432"
        }
	log.Println(serverConfig)
}
