package seeds

import (
	"log"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/gorm"
)

// SeedCountries creates default countries
func SeedCountries(db *gorm.DB) map[string]uuid.UUID {
	var count int64
	db.Model(&models.Country{}).Count(&count)
	if count > 0 {
		log.Printf("📦 Countries table already has %d records, skipping seeding.\n", count)
		return nil
	}

	countries := []models.Country{
		{Name: "Germany", Code: "DE"},
		{Name: "France", Code: "FR"},
		{Name: "Italy", Code: "IT"},
	}

	// Генерация UUID для каждой страны
	for i := range countries {
		countries[i].ID, _ = uuid.NewV4()
	}

	if err := db.Create(&countries).Error; err != nil {
		log.Fatalf("❌ Failed to seed countries: %v", err)
	}

	result := make(map[string]uuid.UUID)
	for _, c := range countries {
		result[c.Name] = c.ID
	}
	log.Println("✅ Countries seeded successfully!")
	return result
}
