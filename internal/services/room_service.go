package services

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	"github.com/vitali-q/selena-hotels-service/internal/models"
)

var (
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

type RoomService struct {
	repo RoomRepository
}

func NewRoomService(repo RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom(room *models.Room) (*models.Room, error) {
	if err := validateRoom(room); err != nil {
		return nil, err
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
