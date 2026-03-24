package config

import (
	"os"
)

// Config содержит все настройки приложения
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

// Load читает конфигурацию из env переменных
func Load() *Config {
	return &Config{
		Env:       os.Getenv("PROJECT_SUFFIX"),

		Port:       os.Getenv("HOTELS_SERVICE_PORT"),

		DBHost:     os.Getenv("HOTELS_COCKROACH_HOST"),
		DBUser:     os.Getenv("HOTELS_COCKROACH_USER"),
		DBPassword: os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		DBName:     os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		DBPort:     os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),

		ProjectSuffix: os.Getenv("PROJECT_SUFFIX"),
	}
}