package bootstrap

import (
	"fmt"
	"log"

	"github.com/vitali-q/hotels-service/internal/config"
	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/handlers"
	"github.com/vitali-q/hotels-service/internal/repository"
	"github.com/vitali-q/hotels-service/internal/services"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB              *gorm.DB
	HotelHandler    *handlers.HotelHandler
	LocationHandler *handlers.LocationHandler
}

func Init() (*Dependencies, error) {
	// --- Database ---
	log.Println("🌱 Initializing database...")
	db, err := database.Init(config.Load())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	log.Println("✅ Database initialized")

	// --- Repositories ---
	hotelRepo := repository.NewHotelRepository(db)

	// --- Services ---
	hotelService := services.NewHotelService(hotelRepo)
	locationService := services.NewLocationService(db)

	// --- Handlers ---
	hotelHandler := handlers.NewHotelHandler(hotelService)
	locationHandler := handlers.NewLocationHandler(locationService)

	return &Dependencies{
		DB:              db,
		HotelHandler:    hotelHandler,
		LocationHandler: locationHandler,
	}, nil
}