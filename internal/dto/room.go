package dto

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	"github.com/vitali-q/selena-hotels-service/internal/models"
)

type CreateRoomRequest struct {
	Number        string          `json:"number" binding:"required,max=50"`
	Type          string          `json:"type" binding:"required,max=50"`
	Capacity      int             `json:"capacity" binding:"required,gt=0"`
	PricePerNight decimal.Decimal `json:"price_per_night" binding:"required"`
}

type UpdateRoomRequest struct {
	Number        *string          `json:"number" binding:"omitempty,max=50"`
	Type          *string          `json:"type" binding:"omitempty,max=50"`
	Capacity      *int             `json:"capacity" binding:"omitempty,gt=0"`
	PricePerNight *decimal.Decimal `json:"price_per_night"`
	Active        *bool            `json:"active"`
}

type RoomResponse struct {
	ID            uuid.UUID       `json:"id"`
	HotelID       uuid.UUID       `json:"hotel_id"`
	Number        string          `json:"number"`
	Type          string          `json:"type"`
	Capacity      int             `json:"capacity"`
	PricePerNight decimal.Decimal `json:"price_per_night"`
	Active        bool            `json:"active"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

func MapRoom(room *models.Room) RoomResponse {
	return RoomResponse{ID: room.ID, HotelID: room.HotelID, Number: room.Number, Type: room.Type, Capacity: room.Capacity, PricePerNight: room.PricePerNight, Active: room.Active, CreatedAt: room.CreatedAt, UpdatedAt: room.UpdatedAt}
}
