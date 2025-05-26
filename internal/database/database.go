package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("HOTELS_COCKROACH_HOST"),
		os.Getenv("HOTELS_COCKROACH_USER"),
		os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
