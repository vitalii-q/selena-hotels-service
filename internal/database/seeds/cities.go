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

	result := make(map[string]uuid.UUID)

	if count > 0 {
		log.Printf("📦 Cities table already has %d records, skipping seeding.\n", count)
		
		// Забираем все города из базы, чтобы заполнить карту
		var existingCities []models.City
		db.Find(&existingCities)
		for _, c := range existingCities {
			result[c.Name] = c.ID
		}
		return result
	}

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

		// France
		{Name: "Paris", CountryID: countries["France"], Capital: true},
		{Name: "Lyon", CountryID: countries["France"]},
		{Name: "Marseille", CountryID: countries["France"]},
		{Name: "Nice", CountryID: countries["France"]},

		// Italy
		{Name: "Rome", CountryID: countries["Italy"], Capital: true},
		{Name: "Milan", CountryID: countries["Italy"]},
		{Name: "Venice", CountryID: countries["Italy"]},
		{Name: "Florence", CountryID: countries["Italy"]},

		// United Kingdom
		{Name: "London", CountryID: countries["United Kingdom"], Capital: true},
		{Name: "Manchester", CountryID: countries["United Kingdom"]},
		{Name: "Birmingham", CountryID: countries["United Kingdom"]},

		// Spain
		{Name: "Madrid", CountryID: countries["Spain"], Capital: true},
		{Name: "Barcelona", CountryID: countries["Spain"]},
		{Name: "Valencia", CountryID: countries["Spain"]},

		// Poland
		{Name: "Warsaw", CountryID: countries["Poland"], Capital: true},
		{Name: "Krakow", CountryID: countries["Poland"]},
		{Name: "Gdansk", CountryID: countries["Poland"]},

		// Switzerland
		{Name: "Bern", CountryID: countries["Switzerland"], Capital: true},
		{Name: "Zurich", CountryID: countries["Switzerland"]},
		{Name: "Geneva", CountryID: countries["Switzerland"]},

		// Portugal
		{Name: "Lisbon", CountryID: countries["Portugal"], Capital: true},
		{Name: "Porto", CountryID: countries["Portugal"]},
		{Name: "Funchal", CountryID: countries["Portugal"]},

		// Netherlands
		{Name: "Amsterdam", CountryID: countries["Netherlands"], Capital: true},
		{Name: "Rotterdam", CountryID: countries["Netherlands"]},
		{Name: "The Hague", CountryID: countries["Netherlands"]},

		// Belgium
		{Name: "Brussels", CountryID: countries["Belgium"], Capital: true},
		{Name: "Antwerp", CountryID: countries["Belgium"]},
		{Name: "Ghent", CountryID: countries["Belgium"]},

		// Australia
		{Name: "Canberra", CountryID: countries["Australia"], Capital: true},
		{Name: "Sydney", CountryID: countries["Australia"]},
		{Name: "Melbourne", CountryID: countries["Australia"]},

		// Finland
		{Name: "Helsinki", CountryID: countries["Finland"], Capital: true},
		{Name: "Espoo", CountryID: countries["Finland"]},
		{Name: "Tampere", CountryID: countries["Finland"]},

		// Ukraine
		{Name: "Kyiv", CountryID: countries["Ukraine"], Capital: true},
		{Name: "Lviv", CountryID: countries["Ukraine"]},
		{Name: "Odessa", CountryID: countries["Ukraine"]},

		// United States
		{Name: "Washington D.C.", CountryID: countries["United States"], Capital: true},
		{Name: "New York", CountryID: countries["United States"]},
		{Name: "Los Angeles", CountryID: countries["United States"]},
		{Name: "Chicago", CountryID: countries["United States"]},

		// Canada
		{Name: "Ottawa", CountryID: countries["Canada"], Capital: true},
		{Name: "Toronto", CountryID: countries["Canada"]},
		{Name: "Vancouver", CountryID: countries["Canada"]},

		// China
		{Name: "Beijing", CountryID: countries["China"], Capital: true},
		{Name: "Shanghai", CountryID: countries["China"]},
		{Name: "Shenzhen", CountryID: countries["China"]},

		// Japan
		{Name: "Tokyo", CountryID: countries["Japan"], Capital: true},
		{Name: "Osaka", CountryID: countries["Japan"]},
		{Name: "Kyoto", CountryID: countries["Japan"]},
	}

	for i := range cities {
		cities[i].ID, _ = uuid.NewV4()
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatalf("❌ Failed to seed cities: %v", err)
	}

	for _, c := range cities {
		result[c.Name] = c.ID
	}

	log.Println("✅ Cities seeded successfully!")
	return result
}
