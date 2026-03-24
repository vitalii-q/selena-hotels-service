package services

import (
	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/dto"
	"github.com/vitali-q/hotels-service/internal/models"
	"gorm.io/gorm"
)

type HotelService struct {
	db *gorm.DB
}

func NewHotelService(db *gorm.DB) *HotelService {
	return &HotelService{db: db}
}

func (s *HotelService) CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
    if err := s.db.Create(hotel).Error; err != nil {
        return nil, err
    }
    return hotel, nil
}

func (s *HotelService) GetAllHotels() ([]dto.HotelResponse, error) {
	var hotels []models.Hotel

	if err := s.db.
		Preload("City").
		Preload("Country").
		Find(&hotels).Error; err != nil {
		return nil, err
	}

	result := make([]dto.HotelResponse, 0, len(hotels))
	for _, hotel := range hotels {
		result = append(result, mapHotelToDTO(hotel))
	}

	return result, nil
}


func (s *HotelService) GetHotelByID(id uuid.UUID) (*models.Hotel, error) {
    var hotel models.Hotel
    if err := s.db.First(&hotel, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &hotel, nil
}

func (s *HotelService) UpdateHotel(id uuid.UUID, newHotel *models.Hotel) (*models.Hotel, error) {
    var hotel models.Hotel
    if err := s.db.First(&hotel, "id = ?", id).Error; err != nil {
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

    if err := s.db.Save(&hotel).Error; err != nil {
        return nil, err
    }

    return &hotel, nil
}

func (s *HotelService) DeleteHotel(id uuid.UUID) error {
    if err := s.db.Delete(&models.Hotel{}, "id = ?", id).Error; err != nil {
        return err
    }
    return nil
}

