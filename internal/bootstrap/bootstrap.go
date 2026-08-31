package bootstrap

import (
	"fmt"
	"log"

	"github.com/vitali-q/selena-hotels-service/internal/config"
	"github.com/vitali-q/selena-hotels-service/internal/database"
	"github.com/vitali-q/selena-hotels-service/internal/handlers"
	"github.com/vitali-q/selena-hotels-service/internal/repository"
	"github.com/vitali-q/selena-hotels-service/internal/services"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB              *gorm.DB
	Env             *config.Env
	HotelHandler    *handlers.HotelHandler
	LocationHandler *handlers.LocationHandler
	RoomHandler     *handlers.RoomHandler
}

func Init() (*Dependencies, error) {
	// --- Configs from .env file ---
	env := config.LoadEnv()

	// --- Database ---
	log.Println("🌱 Initializing database...")
	db, err := database.Init(env)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	log.Println("✅ Database initialized")

	// --- Repositories ---
	hotelRepo := repository.NewHotelRepository(db)
	roomRepo := repository.NewRoomRepository(db)

	// --- Services ---
	hotelService := services.NewHotelService(hotelRepo)
	roomService := services.NewRoomService(roomRepo, hotelRepo, roomRepo)
	locationService := services.NewLocationService(db)

	// --- Handlers ---
	hotelHandler := handlers.NewHotelHandler(hotelService)
	locationHandler := handlers.NewLocationHandler(locationService)
	roomHandler := handlers.NewRoomHandler(roomService)

	return &Dependencies{
		DB:              db,
		Env:             env,
		HotelHandler:    hotelHandler,
		LocationHandler: locationHandler,
		RoomHandler:     roomHandler,
	}, nil
}
