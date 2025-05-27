package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	user := os.Getenv("HOTELS_COCKROACH_USER")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s sslcert=%s sslkey=%s",
		os.Getenv("HOTELS_COCKROACH_HOST"),
		user,
		os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		os.Getenv("DB_SSLMODE"),
		"/certs/ca.crt",
		fmt.Sprintf("/certs/client.%s.crt", user),
		fmt.Sprintf("/certs/client.%s.key", user),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

