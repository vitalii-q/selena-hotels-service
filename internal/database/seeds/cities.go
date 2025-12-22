package seeds

import (
	"log"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/gorm"
)

// SeedCities creates default cities linked to existing countries
func SeedCities(db *gorm.DB, countries map[string]uuid.UUID) map[string]uuid.UUID {
	var count int64
	db.Model(&models.City{}).Count(&count)
	if count > 0 {
		log.Printf("📦 Cities table already has %d records, skipping seeding.\n", count)
		return nil
	}

	// Список городов с привязкой к стране
	cities := []models.City{
		{Name: "Berlin", CountryID: countries["Germany"]},
		{Name: "Hamburg", CountryID: countries["Germany"]},
		{Name: "Munich", CountryID: countries["Germany"]},
		{Name: "Frankfurt", CountryID: countries["Germany"]},
		{Name: "Cologne", CountryID: countries["Germany"]},
		{Name: "Stuttgart", CountryID: countries["Germany"]},
		{Name: "Dresden", CountryID: countries["Germany"]},
		{Name: "Garmisch-Partenkirchen", CountryID: countries["Germany"]},
		{Name: "Heidelberg", CountryID: countries["Germany"]},
		{Name: "Baden-Baden", CountryID: countries["Germany"]},
		{Name: "Augsburg", CountryID: countries["Germany"]},
	}

	for i := range cities {
		cities[i].ID, _ = uuid.NewV4()
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatalf("❌ Failed to seed cities: %v", err)
	}

	result := make(map[string]uuid.UUID)
	for _, c := range cities {
		result[c.Name] = c.ID
	}
	log.Println("✅ Cities seeded successfully!")
	return result
}
