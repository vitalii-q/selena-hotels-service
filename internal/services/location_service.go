package services

import (
	"github.com/vitali-q/selena-hotels-service/internal/dto"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"gorm.io/gorm"
)

type LocationService struct {
	db *gorm.DB
}

func NewLocationService(db *gorm.DB) *LocationService {
	return &LocationService{db: db}
}

// GetCountriesWithCities returns countries with their cities
func (s *LocationService) GetCountriesWithCities() ([]dto.CountryWithCitiesDTO, error) {
	var countries []models.Country

	// Загружаем страны + связанные города
	if err := s.db.Preload("Cities").Find(&countries).Error; err != nil {
		return nil, err
	}

	result := make([]dto.CountryWithCitiesDTO, 0, len(countries))

	for _, country := range countries {
		cities := make([]dto.CityDTO, 0, len(country.Cities))

		for _, city := range country.Cities {
			cities = append(cities, dto.CityDTO{
				ID:   city.ID,
				Name: city.Name,
			})
		}

		result = append(result, dto.CountryWithCitiesDTO{
			ID:     country.ID,
			Name:   country.Name,
			Code:   country.Code,
			Cities: cities,
		})
	}

	return result, nil
}
