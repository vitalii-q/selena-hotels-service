package services

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"github.com/vitali-q/selena-hotels-service/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrReservationBookingRequired = errors.New("booking ID must be greater than zero")
	ErrReservationRoomRequired    = errors.New("room ID is required")
	ErrReservationDatesInvalid    = errors.New("check-out date must be after check-in date")
	ErrReservationNotAvailable    = repository.ErrRoomReservationConflict
	ErrReservationNotFound        = errors.New("reservation not found")
)

type RoomReservationService struct {
	repository *repository.RoomReservationRepository
}

func NewRoomReservationService(repository *repository.RoomReservationRepository) *RoomReservationService {
	return &RoomReservationService{repository: repository}
}

func (s *RoomReservationService) CreateReservation(roomID uuid.UUID, bookingID int64, checkInDate, checkOutDate time.Time, expiresAt *time.Time) (*models.RoomReservation, bool, error) {
	if roomID == uuid.Nil {
		return nil, false, ErrReservationRoomRequired
	}
	if bookingID <= 0 {
		return nil, false, ErrReservationBookingRequired
	}
	if !checkOutDate.After(checkInDate) {
		return nil, false, ErrReservationDatesInvalid
	}

	reservation := &models.RoomReservation{
		RoomID: roomID, BookingID: bookingID, CheckInDate: checkInDate,
		CheckOutDate: checkOutDate, Status: models.RoomReservationStatusActive,
		ExpiresAt: expiresAt,
	}
	return s.repository.CreateReservation(reservation)
}

func (s *RoomReservationService) ReleaseReservation(bookingID int64) (*models.RoomReservation, error) {
	if bookingID <= 0 {
		return nil, ErrReservationBookingRequired
	}
	reservation, err := s.repository.ReleaseByBookingID(bookingID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrReservationNotFound
	}
	return reservation, err
}
