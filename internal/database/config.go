package database

import (
    "github.com/vitali-q/hotels-service/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func InitDB() {
    dsn := os.Getenv("DATABASE_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    err = db.AutoMigrate(&models.Hotel{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }

    DB = db
}
