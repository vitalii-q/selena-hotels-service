package seeds

import (
	"log"

	"github.com/vitali-q/hotels-service/internal/database"
)

// RunSeeds launches all seeds for hotels-service:
// docker exec -it hotels-service go run cmd/seed/main.go
//
// The order of seeding: hotels, locations (hotels-service) -> users (users-service) -> bookings (bookings-service)
func RunSeeds() {
	log.Println("🌱 Initializing database connection...")
	if err := database.Init(); err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}

	log.Println("🌱 Seeding countries...")
	countries := SeedCountries(database.DB)
	//log.Printf("🌱 Countries map: %+v\n", countries)

	log.Println("🌱 Seeding cities...")
	cities := SeedCities(database.DB, countries)

	log.Println("🌱 Starting hotel seeds...")
	if err := SeedHotels(database.DB, cities, countries); err != nil {
		log.Fatalf("❌ Failed to seed hotels: %v", err)
	}

	log.Println("✅ Hotel seeding completed successfully!")
}
