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
	db, err := database.Init(); if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}

	log.Println("🌱 Seeding countries...")
	countries := SeedCountries(db)
	//log.Printf("🌱 Countries map: %+v\n", countries)

	log.Println("🌱 Seeding cities...")
	cities := SeedCities(db, countries)

	log.Println("🌱 Starting hotel seeds...")
	if err := SeedHotels(db, cities, countries); err != nil {
		log.Fatalf("❌ Failed to seed hotels: %v", err)
	}

	log.Println("✅ Hotel seeding completed successfully!")
}
