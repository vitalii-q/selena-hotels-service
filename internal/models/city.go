package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type City struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"not null"`
	CountryID uuid.UUID `json:"country_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *City) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID, err = uuid.NewV4()
	return
}
