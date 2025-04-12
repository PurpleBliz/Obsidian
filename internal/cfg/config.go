package cfg

import (
	"os"
)

type Config struct {
	DBConnectionString string
	AppPort            string
}

func NewConfig() *Config {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		//panic("DATABASE_URL environment variable not set")
		connStr = "postgres://test_user:user_password@obsidian_db:5432/homebd"
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	return &Config{
		DBConnectionString: connStr,
		AppPort:            appPort,
	}
}
