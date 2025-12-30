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

func SeedHotels(db *gorm.DB, cities map[string]uuid.UUID, countries map[string]uuid.UUID) error {
	var count int64
	db.Model(&models.Hotel{}).Count(&count)

	if count > 0 {
		log.Printf("📦 Hotels table already has %d records, skipping seeding.\n", count)
		return nil
	}

	// локальный RNG (Go 1.20+ корректно)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	bigCountries := map[string]bool{
		"Germany": true,
		"France":  true,
		"Italy":   true,
		"USA":     true,
		"China":   true,
		"Japan":   true,
	}

	var hotels []models.Hotel

	for _, cityID := range cities {
		var city models.City
		if err := db.First(&city, "id = ?", cityID).Error; err != nil {
			continue
		}

		var country models.Country
		if err := db.First(&country, "id = ?", city.CountryID).Error; err != nil {
			continue
		}

		minHotels := 2
		maxHotels := 3

		if bigCountries[country.Name] {
			minHotels = 4
			maxHotels = 6
		}

		hotelCount := rng.Intn(maxHotels-minHotels+1) + minHotels

		for i := 0; i < hotelCount; i++ {
			hotels = append(hotels, models.Hotel{
				ID:          newUUID(),
				Name:        strPtr(fmt.Sprintf("%s %s Hotel", city.Name, randomHotelSuffix(rng))),
				Description: strPtr(fmt.Sprintf("Comfortable hotel in %s.", city.Name)),
				Address:     strPtr(generateStreetAddress(rng)),
				CityID:      city.ID,
				CountryID:   country.ID,
				Price:       floatPtr(randomPrice(rng)),
				Amenities:   randomAmenitiesJSON(rng),
			})
		}
	}

	if err := db.Create(&hotels).Error; err != nil {
		return err
	}

	log.Printf("✅ Seeded %d hotels successfully\n", len(hotels))
	return nil
}

/* ---------------- HELPERS ---------------- */

func newUUID() uuid.UUID {
	id, _ := uuid.NewV4()
	return id
}

func strPtr(s string) *string {
	return &s
}

func floatPtr(f float64) *float64 {
	return &f
}

/* -------- STREET GENERATOR -------- */

func generateStreetAddress(rng *rand.Rand) string {
	adjectives := []string{
		"Central", "Old", "New", "Royal", "Green", "Grand", "Sunny",
	}

	objects := []string{
		"Park", "River", "Market", "Garden", "Hill", "Square", "Lake",
	}

	streetTypes := []string{
		"Street", "Avenue", "Road", "Boulevard", "Lane",
	}

	number := rng.Intn(90) + 1

	return fmt.Sprintf(
		"%s %s %s %d",
		randomFrom(rng, adjectives),
		randomFrom(rng, objects),
		randomFrom(rng, streetTypes),
		number,
	)
}

/* -------- AMENITIES -------- */

func randomAmenitiesJSON(rng *rand.Rand) datatypes.JSON {
	allAmenities := []string{
		"WiFi",
		"Breakfast",
		"Parking",
		"Gym",
		"Spa",
		"Pool",
		"Air Conditioning",
		"Pet Friendly",
		"Restaurant",
		"Bar",
	}

	rng.Shuffle(len(allAmenities), func(i, j int) {
		allAmenities[i], allAmenities[j] = allAmenities[j], allAmenities[i]
	})

	count := rng.Intn(2) + 2 // 2–3 amenities
	selected := allAmenities[:count]

	json := fmt.Sprintf(`["%s"]`, joinWithQuotes(selected))
	return datatypes.JSON([]byte(json))
}

/* -------- MISC -------- */

func randomFrom(rng *rand.Rand, list []string) string {
	return list[rng.Intn(len(list))]
}

func joinWithQuotes(items []string) string {
	result := ""
	for i, v := range items {
		if i > 0 {
			result += `","`
		}
		result += v
	}
	return result
}

func randomHotelSuffix(rng *rand.Rand) string {
	suffixes := []string{
		"Central", "Grand", "Plaza", "Inn", "Suites", "Residence",
	}
	return randomFrom(rng, suffixes)
}

func randomPrice(rng *rand.Rand) float64 {
	return float64(rng.Intn(180) + 70) // 70–250
}
