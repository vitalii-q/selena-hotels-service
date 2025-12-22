package seeds

import (
	"log"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/gorm"
)

// SeedCountries creates default countries
func SeedCountries(db *gorm.DB) map[string]uuid.UUID {
	var countries []models.Country
	db.Find(&countries) // вытаскивает все существующие страны

	result := make(map[string]uuid.UUID)
	for _, c := range countries {
		result[c.Name] = c.ID
	}

	if len(countries) > 0 {
		log.Printf("📦 Countries table already has %d records, skipping creation.\n", len(countries))
		return result
	}

	// Создаём страны, если их нет
	newCountries := []models.Country{
		{Name: "Germany", Code: "DE"},
		{Name: "France", Code: "FR"},
		{Name: "Italy", Code: "IT"},
	}

	for i := range newCountries {
		newCountries[i].ID, _ = uuid.NewV4()
	}

	if err := db.Create(&newCountries).Error; err != nil {
		log.Fatalf("❌ Failed to seed countries: %v", err)
	}

	for _, c := range newCountries {
		result[c.Name] = c.ID
	}

	log.Println("✅ Countries seeded successfully!")
	return result
}

