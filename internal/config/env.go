package config

import (
	"os"
	"strconv"
	"time"
)

// Config contains all the application settings
type Env struct {
	AppEnv string

	Port string

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string

	ProjectSuffix             string
	ReservationExpiryEnabled  bool
	ReservationExpiryInterval time.Duration
}

// Load reads configuration from env variables
func LoadEnv() *Env {
	enabled := true
	if value := os.Getenv("HOTELS_RESERVATION_EXPIRY_ENABLED"); value != "" {
		enabled, _ = strconv.ParseBool(value)
	}
	interval := time.Minute
	if value := os.Getenv("HOTELS_RESERVATION_EXPIRY_INTERVAL"); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			interval = parsed
		}
	}
	env := &Env{
		AppEnv: os.Getenv("APP_ENV"),

		Port: os.Getenv("HOTELS_SERVICE_PORT"),

		DBHost:                    os.Getenv("HOTELS_COCKROACH_HOST"),
		DBUser:                    os.Getenv("HOTELS_COCKROACH_USER"),
		DBPassword:                os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		DBName:                    os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		DBPort:                    os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		DBSSLMode:                 os.Getenv("DB_SSLMODE"),
		ReservationExpiryEnabled:  enabled,
		ReservationExpiryInterval: interval,
	}

	// --- Validate required configuration ---
	if env.Port == "" {
		panic("HOTELS_SERVICE_PORT is not set")
	}

	return env
}
