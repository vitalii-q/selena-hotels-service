package config

import (
	"os"
)

// Config contains all the application settings
type Config struct {
	Env string

	Port string

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string

	ProjectSuffix string
}

// Load reads configuration from env variables
func Load() *Config {
	cfg := &Config{
		Env:        os.Getenv("PROJECT_SUFFIX"), // TODO: Remane to APP_ENV

		Port:       os.Getenv("HOTELS_SERVICE_PORT"),

		DBHost:     os.Getenv("HOTELS_COCKROACH_HOST"),
		DBUser:     os.Getenv("HOTELS_COCKROACH_USER"),
		DBPassword: os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		DBName:     os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		DBPort:     os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// --- Validate required configuration --- 
	if cfg.Port == "" {
		panic("HOTELS_SERVICE_PORT is not set")
	}

	return cfg
}