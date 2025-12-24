package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
    "gorm.io/datatypes"
)

type Hotel struct {
    ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
    Name        *string        `json:"name,omitempty"`
    Description *string        `json:"description,omitempty"`
    Address     *string        `json:"address,omitempty"`
	CityID      uuid.UUID      `json:"city_id"`
	CountryID   uuid.UUID      `json:"country_id"`
    Price       *float64       `json:"price,omitempty"`
	Amenities   datatypes.JSON `json:"amenities" gorm:"type:jsonb"`

	City        City           `gorm:"foreignKey:CityID"`
	Country     Country        `gorm:"foreignKey:CountryID"`

    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
}

func (h *Hotel) BeforeCreate(tx *gorm.DB) (err error) {  // a function that is automatically invoked before creation
	h.ID, err = uuid.NewV4()                             // uuid generation
	return
}
