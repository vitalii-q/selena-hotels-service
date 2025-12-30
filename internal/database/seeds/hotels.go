package seeds

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// SeedHotels creates hotels for all cities from database
func SeedHotels(db *gorm.DB, cities map[string]uuid.UUID, countries map[string]uuid.UUID) error {
	var count int64
	db.Model(&models.Hotel{}).Count(&count)
	if count > 0 {
		log.Printf("📦 Hotels table already has %d records, skipping seeding.\n", count)
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	// Загружаем города вместе со странами
	var dbCities []models.City
	if err := db.Preload("Country").Find(&dbCities).Error; err != nil {
		return err
	}

	// Страны с большим количеством отелей
	bigCountries := map[string]bool{
		"Germany": true,
		"France":  true,
		"United States": true,
		"China":   true,
		"Japan":   true,
		"Italy":   true,
	}

	var hotels []models.Hotel

	for _, city := range dbCities {
		var minHotels, maxHotels int

		if bigCountries[city.Country.Name] {
			minHotels = 4
			maxHotels = 6
		} else {
			minHotels = 2
			maxHotels = 3
		}

		hotelsCount := rand.Intn(maxHotels-minHotels+1) + minHotels

		for i := 1; i <= hotelsCount; i++ {
			hotel := models.Hotel{
				ID:          uuid.Must(uuid.NewV4()),
				Name:        strPtr(fmt.Sprintf("%s Hotel %d", city.Name, i)),
				Description: strPtr(fmt.Sprintf("Comfortable hotel in %s.", city.Name)),
				Address:     strPtr(fmt.Sprintf("Main street %d, %s", i*3, city.Name)),
				CityID:      city.ID,
				CountryID:   city.CountryID,
				Price:       floatPtr(float64(rand.Intn(120) + 80)),
				Amenities:   datatypes.JSON([]byte(`["WiFi","Breakfast","Parking"]`)),
			}

			hotels = append(hotels, hotel)
		}
	}

	if err := db.Create(&hotels).Error; err != nil {
		return err
	}

	log.Printf("✅ Seeded %d hotels successfully!\n", len(hotels))
	return nil
}

// helpers
func strPtr(s string) *string {
	return &s
}

func floatPtr(f float64) *float64 {
	return &f
}
