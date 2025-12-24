package services

import (
	"github.com/vitali-q/hotels-service/internal/dto"
	"github.com/vitali-q/hotels-service/internal/models"
)

func mapHotelToDTO(hotel models.Hotel) dto.HotelResponse {
	return dto.HotelResponse{
		ID:      hotel.ID,
		Name:    *hotel.Name,
		Address: *hotel.Address,
		Price:   *hotel.Price,
		City: dto.CityDTO{
			ID:   hotel.City.ID,
			Name: hotel.City.Name,
		},
		Country: dto.CountryDTO{
			ID:   hotel.Country.ID,
			Name: hotel.Country.Name,
			Code: hotel.Country.Code,
		},
	}
}
