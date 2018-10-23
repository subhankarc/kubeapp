package models

type Config struct {
	DB      DBConfig `json:"db"`
	AppPort string   `json:"app_port"`
}

type DBConfig struct {
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	DBName     string `json:"dbname"`
}
