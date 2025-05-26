package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
}

func (h *Hotel) BeforeCreate(tx *gorm.DB) (err error) { // автовызываемая перед созданием функция
	h.ID, err = uuid.NewV4() // генерация uuid
	return
}
