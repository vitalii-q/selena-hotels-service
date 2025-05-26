package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
    ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
    Name        *string   `json:"name,omitempty"`
    Description *string   `json:"description,omitempty"`
    Address     *string   `json:"address,omitempty"`
    City        *string   `json:"city,omitempty"`
    Country     *string   `json:"country,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (h *Hotel) BeforeCreate(tx *gorm.DB) (err error) { // автовызываемая перед созданием функция
	h.ID, err = uuid.NewV4() // генерация uuid
	return
}
