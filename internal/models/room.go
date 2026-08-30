package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Room struct {
	ID            uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	HotelID       uuid.UUID       `json:"hotel_id" gorm:"type:uuid;not null;index"`
	Number        string          `json:"number" gorm:"size:50;not null"`
	Type          string          `json:"type" gorm:"size:50;not null"`
	Capacity      int             `json:"capacity" gorm:"not null"`
	PricePerNight decimal.Decimal `json:"price_per_night" gorm:"type:decimal(10,2);not null"`
	Active        bool            `json:"active" gorm:"not null;default:true"`

	Hotel Hotel `json:"-" gorm:"foreignKey:HotelID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *Room) BeforeCreate(tx *gorm.DB) error {
	if r.ID != uuid.Nil {
		return nil
	}

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	r.ID = id
	return nil
}
