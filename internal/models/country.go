package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Country struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Code      string    `json:"code" gorm:"unique;not null"`

	Cities    []City    `gorm:"foreignKey:CountryID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Country) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID, err = uuid.NewV4()
	return
}
