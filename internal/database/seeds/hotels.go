package seeds

import (
	"log"

	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// SeedHotels creates test hotel records if the table is empty
// docker exec -it hotels-service go run cmd/seed/main.go
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
			Price:       floatPtr(120.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Parking","Rooftop Bar"]`)),
		},
		{
			Name:        strPtr("Sea Breeze Resort"),
			Description: strPtr("Cozy seaside resort with ocean view."),
			Address:     strPtr("Coast Road 8"),
			City:        strPtr("Hamburg"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(180.50),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Pool","Ocean View"]`)),
		},
		{
			Name:        strPtr("Mountain View Inn"),
			Description: strPtr("Quiet retreat near the Alps."),
			Address:     strPtr("Bergstraße 5"),
			City:        strPtr("Munich"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(95.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Free Breakfast","Mountain View"]`)),
		},
		{
			Name:        strPtr("Sunset Lodge"),
			Description: strPtr("Charming lodge with panoramic sunset views."),
			Address:     strPtr("Sunset Blvd 45"),
			City:        strPtr("Frankfurt"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(210.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Spa","Gym","Conference Rooms"]`)),
		},
		{
			Name:        strPtr("City Center Hotel"),
			Description: strPtr("Located in the heart of the city, perfect for business trips."),
			Address:     strPtr("Central Avenue 10"),
			City:        strPtr("Cologne"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(150.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Gym","Sunset View"]`)),
		},
		{
			Name:        strPtr("Lakeside Inn"),
			Description: strPtr("Peaceful retreat next to the lake."),
			Address:     strPtr("Lake Road 3"),
			City:        strPtr("Stuttgart"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(130.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Lake View","Breakfast Included"]`)),
		},
		{
			Name:        strPtr("Historic Grand Hotel"),
			Description: strPtr("Luxury hotel with historic architecture."),
			Address:     strPtr("Old Town 7"),
			City:        strPtr("Dresden"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(200.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Spa","Historic Architecture"]`)),
		},
		{
			Name:        strPtr("Alpine Escape"),
			Description: strPtr("Secluded cabin resort in the Alps."),
			Address:     strPtr("Alpenweg 15"),
			City:        strPtr("Garmisch-Partenkirchen"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(180.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Mountain View","Sauna"]`)),
		},
		{
			Name:        strPtr("Riverside Retreat"),
			Description: strPtr("Quiet retreat by the river with modern amenities."),
			Address:     strPtr("River Road 22"),
			City:        strPtr("Heidelberg"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(140.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","River View","Gym"]`)),
		},
		{
			Name:        strPtr("Forest Haven"),
			Description: strPtr("Cozy cabins surrounded by forest nature."),
			Address:     strPtr("Forest Lane 9"),
			City:        strPtr("Baden-Baden"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(125.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Forest View","Hiking Trails"]`)),
		},
		{
			Name:        strPtr("Augsburg Grand Hotel"),
			Description: strPtr("Luxury hotel in the city center with modern amenities."),
			Address:     strPtr("Maximilianstraße 15"),
			City:        strPtr("Augsburg"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(210.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","Spa","Gym","Conference Rooms"]`)),
		},
		{
			Name:        strPtr("Lech Riverside Inn"),
			Description: strPtr("Comfortable stay along the Lech river with cozy rooms."),
			Address:     strPtr("Lechstraße 7"),
			City:        strPtr("Augsburg"),
			Country:     strPtr("Germany"),
			Price:       floatPtr(145.00),
			Amenities:   datatypes.JSON([]byte(`["WiFi","River View","Breakfast Included"]`)),
		},
	}

	if err := db.Create(&hotels).Error; err != nil {
		return err
	}

	return nil
}

// вспомогательная функция для *string
func strPtr(s string) *string {
	return &s
}

func floatPtr(f float64) *float64 {
    return &f
}