package seeds

import (
	"log"

	"github.com/vitali-q/hotels-service/internal/database"
)

// RunSeeds запускает все сиды для hotels-service
// docker exec -it users-service_dev go run cmd/seed/main.go
func RunSeeds() {
	log.Println("🌱 Initializing database connection...")
	if err := database.Init(); err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}

	log.Println("🌱 Starting hotel seeds...")
	if err := SeedHotels(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed hotels: %v", err)
	}

	log.Println("✅ Hotel seeding completed successfully!")
}
