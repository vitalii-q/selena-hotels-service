package repository

import (
	"github.com/gofrs/uuid"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{db: db}
}

// CreateHotel saves a new hotel
func (r *HotelRepository) CreateHotel(hotel *models.Hotel) error {
	return r.db.Create(hotel).Error
}

// GetAllHotels returns all hotels with preloaded city and country
func (r *HotelRepository) GetAllHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	err := r.db.Preload("City").Preload("Country").Find(&hotels).Error
	return hotels, err
}

// GetHotelByID finds hotel by ID
func (r *HotelRepository) GetHotelByID(id uuid.UUID) (*models.Hotel, error) {
	var hotel models.Hotel
	err := r.db.First(&hotel, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (r *HotelRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Hotel{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpdateHotel updates a hotel
func (r *HotelRepository) UpdateHotel(hotel *models.Hotel) error {
	return r.db.Save(hotel).Error
}

// DeleteHotel deletes a hotel
func (r *HotelRepository) DeleteHotel(id uuid.UUID) error {
	return r.db.Delete(&models.Hotel{}, "id = ?", id).Error
}
