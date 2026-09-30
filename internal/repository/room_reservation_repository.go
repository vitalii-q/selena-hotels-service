package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrRoomReservationConflict = errors.New("room is not available for the requested dates")

type RoomReservationRepository struct {
	db *gorm.DB
}

func NewRoomReservationRepository(db *gorm.DB) *RoomReservationRepository {
	return &RoomReservationRepository{db: db}
}

func (r *RoomReservationRepository) CreateReservation(reservation *models.RoomReservation) (*models.RoomReservation, bool, error) {
	const maxRetries = 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		created, existing, err := r.createReservationOnce(reservation)
		if err == nil {
			return created, existing, nil
		}
		if isUniqueViolation(err) {
			current, findErr := r.findByBookingID(reservation.BookingID)
			if findErr == nil {
				return current, true, nil
			}
		}
		if !isSerializationError(err) || attempt == maxRetries-1 {
			return nil, false, err
		}
		time.Sleep(time.Duration(attempt+1) * 25 * time.Millisecond)
	}
	return nil, false, fmt.Errorf("reservation transaction retry limit exceeded")
}

func (r *RoomReservationRepository) findByBookingID(bookingID int64) (*models.RoomReservation, error) {
	var reservation models.RoomReservation
	err := r.db.Where("booking_id = ?", bookingID).Order("created_at DESC").First(&reservation).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *RoomReservationRepository) createReservationOnce(reservation *models.RoomReservation) (*models.RoomReservation, bool, error) {
	var result *models.RoomReservation
	existing := false
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var current models.RoomReservation
		findErr := tx.Where("booking_id = ?", reservation.BookingID).Order("created_at DESC").First(&current).Error
		if findErr == nil {
			result = &current
			existing = true
			return nil
		}
		if !errors.Is(findErr, gorm.ErrRecordNotFound) {
			return findErr
		}

		var room models.Room
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&room, "id = ?", reservation.RoomID).Error; err != nil {
			return err
		}

		var conflicting models.RoomReservation
		conflictErr := tx.Where("room_id = ? AND status = ?", reservation.RoomID, models.RoomReservationStatusActive).
			Where("check_in_date < ? AND check_out_date > ?", reservation.CheckOutDate, reservation.CheckInDate).
			First(&conflicting).Error
		if conflictErr == nil {
			return ErrRoomReservationConflict
		}
		if !errors.Is(conflictErr, gorm.ErrRecordNotFound) {
			return conflictErr
		}

		if err := tx.Create(reservation).Error; err != nil {
			return err
		}
		result = reservation
		return nil
	})
	return result, existing, err
}

func isSerializationError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "40001" {
		return true
	}
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "restart transaction") || strings.Contains(message, "serialization")
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
