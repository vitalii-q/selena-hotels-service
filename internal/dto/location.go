package dto

import "github.com/gofrs/uuid"

type CityDTO struct {
	ID   uuid.UUID   `json:"id"`
	Name string      `json:"name"`
	Capital bool     `json:"capital"`
}

type CountryWithCitiesDTO struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Cities []CityDTO `json:"cities"`
}
