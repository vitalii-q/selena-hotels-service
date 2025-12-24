package dto

import "github.com/gofrs/uuid"

type HotelResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Price   float64   `json:"price"`

	City    CityDTO    `json:"city"`
	Country CountryDTO `json:"country"`
}

type CountryDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}
