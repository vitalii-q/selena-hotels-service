package services

import (
	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/dto"
	"github.com/vitali-q/hotels-service/internal/models"
	"github.com/vitali-q/hotels-service/internal/repository"
)

type HotelService struct {
	repo *repository.HotelRepository
}

func NewHotelService(repo *repository.HotelRepository) *HotelService {
	return &HotelService{repo: repo}
}

func (s *HotelService) CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
	if err := s.repo.CreateHotel(hotel); err != nil {
		return nil, err
	}
    return hotel, nil
}

func (s *HotelService) GetAllHotels() ([]dto.HotelResponse, error) {
	hotels, err := s.repo.GetAllHotels()
	if err != nil {
		return nil, err
	}

	result := make([]dto.HotelResponse, 0, len(hotels))
	for _, hotel := range hotels {
		result = append(result, mapHotelToDTO(hotel))
	}
	return result, nil
}


func (s *HotelService) GetHotelByID(id uuid.UUID) (*models.Hotel, error) {
	return s.repo.GetHotelByID(id)
}

func (s *HotelService) UpdateHotel(id uuid.UUID, newHotel *models.Hotel) (*models.Hotel, error) {
	hotel, err := s.repo.GetHotelByID(id)
	if err != nil {
		return nil, err
	}

	if newHotel.Name != nil {
		hotel.Name = newHotel.Name
	}
	if newHotel.Description != nil {
		hotel.Description = newHotel.Description
	}
	if newHotel.Address != nil {
		hotel.Address = newHotel.Address
	}
	if newHotel.CityID != uuid.Nil {
		hotel.CityID = newHotel.CityID
	}
	if newHotel.CountryID != uuid.Nil {
		hotel.CountryID = newHotel.CountryID
	}

	if err := s.repo.UpdateHotel(hotel); err != nil {
		return nil, err
	}

	return hotel, nil
}

func (s *HotelService) DeleteHotel(id uuid.UUID) error {
	return s.repo.DeleteHotel(id)
}

