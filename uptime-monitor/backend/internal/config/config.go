package config

import (
	"os"
	"time"
)

type Config struct {
	DBConnectionString string
	Urls               []string
	Interval           time.Duration
}

func LoadConfig() *Config {
	return &Config{
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		Urls:               []string{"https://google.com", "https://github.com"},
		Interval:           1 * time.Minute,
	}
}
