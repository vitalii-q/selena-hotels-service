package models

import "github.com/gofrs/uuid"

type Hotel struct {
    ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Address     string `json:"address"`
    City        string `json:"city"`
    Country     string `json:"country"`
}
