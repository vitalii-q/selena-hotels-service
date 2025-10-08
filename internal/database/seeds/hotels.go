package seeds

import (
	"log"

	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/gorm"
)

// SeedHotels создает тестовые записи отелей, если таблица пуста
func SeedHotels(db *gorm.DB) error {
	var count int64
	db.Model(&models.Hotel{}).Count(&count)

	if count > 0 {
		log.Printf("📦 Hotels table already has %d records, skipping seeding.\n", count)
		return nil
	}

	hotels := []models.Hotel{
		{
			Name:        strPtr("Hotel Aurora"),
			Description: strPtr("A modern 4-star hotel with a rooftop terrace."),
			Address:     strPtr("Main Street 12"),
			City:        strPtr("Berlin"),
			Country:     strPtr("Germany"),
		},
		{
			Name:        strPtr("Sea Breeze Resort"),
			Description: strPtr("Cozy seaside resort with ocean view."),
			Address:     strPtr("Coast Road 8"),
			City:        strPtr("Hamburg"),
			Country:     strPtr("Germany"),
		},
		{
			Name:        strPtr("Mountain View Inn"),
			Description: strPtr("Quiet retreat near the Alps."),
			Address:     strPtr("Bergstraße 5"),
			City:        strPtr("Munich"),
			Country:     strPtr("Germany"),
		},
	}

	if err := db.Create(&hotels).Error; err != nil {
		return err
	}

	log.Println("✅ Hotel seeds inserted successfully.")
	return nil
}

// вспомогательная функция для *string
func strPtr(s string) *string {
	return &s
}
