package services

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	"github.com/vitali-q/selena-hotels-service/internal/models"
)

var (
	ErrRoomNotFound             = errors.New("room not found")
	ErrHotelNotFound            = errors.New("hotel not found")
	ErrActiveReservations       = errors.New("room has active reservations")
	ErrRoomHotelRequired        = errors.New("room hotel ID is required")
	ErrRoomNumberRequired       = errors.New("room number is required")
	ErrRoomTypeRequired         = errors.New("room type is required")
	ErrRoomCapacityInvalid      = errors.New("room capacity must be greater than zero")
	ErrRoomPricePerNightInvalid = errors.New("room price per night must be greater than zero")
)

type RoomRepository interface {
	CreateRoom(room *models.Room) error
	GetRoomByID(id uuid.UUID) (*models.Room, error)
	GetRoomsByHotelID(hotelID uuid.UUID) ([]models.Room, error)
	UpdateRoom(room *models.Room) error
	DeleteRoom(id uuid.UUID) error
}

type HotelLookup interface {
	Exists(id uuid.UUID) (bool, error)
}
type ActiveReservationChecker interface {
	HasActiveReservations(id uuid.UUID) (bool, error)
}

type RoomService struct {
	repo               RoomRepository
	hotelLookup        HotelLookup
	reservationChecker ActiveReservationChecker
}

func NewRoomService(repo RoomRepository, lookups ...interface{}) *RoomService {
	s := &RoomService{repo: repo}
	for _, lookup := range lookups {
		switch v := lookup.(type) {
		case HotelLookup:
			s.hotelLookup = v
		case ActiveReservationChecker:
			s.reservationChecker = v
		}
	}
	return s
}

func (s *RoomService) CreateRoom(room *models.Room) (*models.Room, error) {
	if err := validateRoom(room); err != nil {
		return nil, err
	}
	if s.hotelLookup != nil {
		exists, err := s.hotelLookup.Exists(room.HotelID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, ErrHotelNotFound
		}
	}

	if err := s.repo.CreateRoom(room); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) GetRoomByID(id uuid.UUID) (*models.Room, error) {
	return s.repo.GetRoomByID(id)
}

func (s *RoomService) GetRoomsByHotelID(hotelID uuid.UUID) ([]models.Room, error) {
	return s.repo.GetRoomsByHotelID(hotelID)
}

func (s *RoomService) UpdateRoom(room *models.Room) (*models.Room, error) {
	if err := validateRoom(room); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateRoom(room); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) DeleteRoom(id uuid.UUID) error {
	if s.reservationChecker != nil {
		has, err := s.reservationChecker.HasActiveReservations(id)
		if err != nil {
			return err
		}
		if has {
			return ErrActiveReservations
		}
	}
	return s.repo.DeleteRoom(id)
}

func validateRoom(room *models.Room) error {
	if room == nil || room.HotelID == uuid.Nil {
		return ErrRoomHotelRequired
	}
	if strings.TrimSpace(room.Number) == "" {
		return ErrRoomNumberRequired
	}
	if strings.TrimSpace(room.Type) == "" {
		return ErrRoomTypeRequired
	}
	if room.Capacity <= 0 {
		return ErrRoomCapacityInvalid
	}
	if room.PricePerNight.LessThanOrEqual(decimal.Zero) {
		return ErrRoomPricePerNightInvalid
	}

	return nil
}
