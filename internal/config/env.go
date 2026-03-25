package config

import (
	"os"
)

// Config contains all the application settings
type Env struct {
	AppEnv        string

	Port          string

	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	DBSSLMode     string

	ProjectSuffix string
}

// Load reads configuration from env variables
func Load() *Env {
	env := &Env{
		AppEnv:     os.Getenv("APP_ENV"),

		Port:       os.Getenv("HOTELS_SERVICE_PORT"),

		DBHost:     os.Getenv("HOTELS_COCKROACH_HOST"),
		DBUser:     os.Getenv("HOTELS_COCKROACH_USER"),
		DBPassword: os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		DBName:     os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		DBPort:     os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// --- Validate required configuration --- 
	if env.Port == "" {
		panic("HOTELS_SERVICE_PORT is not set")
	}

	return env
}