package models

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	RoomReservationStatusActive   = "ACTIVE"
	RoomReservationStatusReleased = "RELEASED"
	RoomReservationStatusExpired  = "EXPIRED"
)

var (
	ErrRoomReservationDatesInvalid  = errors.New("room reservation check-out date must be after check-in date")
	ErrRoomReservationStatusInvalid = errors.New("invalid room reservation status")
)

type RoomReservation struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	RoomID       uuid.UUID  `json:"room_id" gorm:"type:uuid;not null;index"`
	BookingID    int64      `json:"booking_id" gorm:"not null;index"`
	CheckInDate  time.Time  `json:"check_in_date" gorm:"type:date;not null"`
	CheckOutDate time.Time  `json:"check_out_date" gorm:"type:date;not null"`
	Status       string     `json:"status" gorm:"size:20;not null;default:ACTIVE"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	Room Room `json:"-" gorm:"foreignKey:RoomID"`
}

func (r *RoomReservation) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		r.ID = id
	}
	return r.Validate()
}

func (r *RoomReservation) BeforeUpdate(tx *gorm.DB) error {
	return r.Validate()
}

func (r *RoomReservation) Validate() error {
	if !r.CheckOutDate.After(r.CheckInDate) {
		return ErrRoomReservationDatesInvalid
	}
	switch r.Status {
	case RoomReservationStatusActive, RoomReservationStatusReleased, RoomReservationStatusExpired:
		return nil
	default:
		return ErrRoomReservationStatusInvalid
	}
}
