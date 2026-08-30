package repository

import (
	"github.com/gofrs/uuid"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) CreateRoom(room *models.Room) error {
	return r.db.Create(room).Error
}

func (r *RoomRepository) GetRoomByID(id uuid.UUID) (*models.Room, error) {
	var room models.Room
	if err := r.db.First(&room, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *RoomRepository) GetRoomsByHotelID(hotelID uuid.UUID) ([]models.Room, error) {
	var rooms []models.Room
	if err := r.db.Where("hotel_id = ?", hotelID).Order("number ASC").Find(&rooms).Error; err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomRepository) UpdateRoom(room *models.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepository) DeleteRoom(id uuid.UUID) error {
	return r.db.Delete(&models.Room{}, "id = ?", id).Error
}
