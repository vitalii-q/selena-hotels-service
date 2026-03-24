package bootstrap

import (
	"fmt"
	"log"

	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/handlers"
	"github.com/vitali-q/hotels-service/internal/repository"
	"github.com/vitali-q/hotels-service/internal/services"
	"gorm.io/gorm"
)

type AppDependencies struct {
	DB             *gorm.DB
	HotelService   *services.HotelService
	HotelHandler   *handlers.HotelHandler
	LocationService *services.LocationService
	LocationHandler *handlers.LocationHandler
}

func Init() (*AppDependencies, error) {
	// --- Database ---
	log.Println("🌱 Initializing database...")
	db, err := database.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	log.Println("✅ Database initialized")

	// --- Hotel DI ---
	hotelRepo := repository.NewHotelRepository(db)
	hotelService := services.NewHotelService(hotelRepo)
	hotelHandler := handlers.NewHotelHandler(hotelService)

	// --- Location DI ---
	locationService := services.NewLocationService(db)
	locationHandler := handlers.NewLocationHandler(locationService)

	return &AppDependencies{
		DB:              db,
		HotelService:    hotelService,
		HotelHandler:    hotelHandler,
		LocationService: locationService,
		LocationHandler: locationHandler,
	}, nil
}