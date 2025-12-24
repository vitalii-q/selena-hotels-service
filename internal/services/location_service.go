package services

import (
	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/dto"
	"github.com/vitali-q/hotels-service/internal/models"
)

// GetCountriesWithCities returns countries with their cities
func GetCountriesWithCities() ([]dto.CountryWithCitiesDTO, error) {
	var countries []models.Country

	// Загружаем страны + связанные города
	if err := database.DB.Preload("Cities").Find(&countries).Error; err != nil {
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
