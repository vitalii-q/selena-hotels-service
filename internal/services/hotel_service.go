package services

import (
	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/models"
)

func CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
    if err := database.DB.Create(hotel).Error; err != nil {
        return nil, err
    }
    return hotel, nil
}

func GetAllHotels() ([]models.Hotel, error) {
    var hotels []models.Hotel
    if err := database.DB.Find(&hotels).Error; err != nil {
        return nil, err
    }
    return hotels, nil
}

func GetHotelByID(id uuid.UUID) (*models.Hotel, error) {
    var hotel models.Hotel
    if err := database.DB.First(&hotel, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &hotel, nil
}

func UpdateHotel(id uuid.UUID, newHotel *models.Hotel) (*models.Hotel, error) {
    var hotel models.Hotel
    if err := database.DB.First(&hotel, "id = ?", id).Error; err != nil {
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
    if newHotel.City != nil {
        hotel.City = newHotel.City
    }
    if newHotel.Country != nil {
        hotel.Country = newHotel.Country
    }

    if err := database.DB.Save(&hotel).Error; err != nil {
        return nil, err
    }

    return &hotel, nil
}

func DeleteHotel(id uuid.UUID) error {
    if err := database.DB.Delete(&models.Hotel{}, "id = ?", id).Error; err != nil {
        return err
    }
    return nil
}

