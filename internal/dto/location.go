package dto

import "github.com/gofrs/uuid"

type CityDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CountryWithCitiesDTO struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Cities []CityDTO `json:"cities"`
}
