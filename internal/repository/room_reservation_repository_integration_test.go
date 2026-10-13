package repository

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vitali-q/selena-hotels-service/internal/config"
	"github.com/vitali-q/selena-hotels-service/internal/database"
	"github.com/vitali-q/selena-hotels-service/internal/models"
)

func TestRoomReservationConcurrentCreateAllowsOneReservation(t *testing.T) {
	if os.Getenv("HOTELS_RUN_INTEGRATION_TESTS") != "1" {
		t.Skip("set HOTELS_RUN_INTEGRATION_TESTS=1 to run against CockroachDB")
	}

	db, err := database.Init(config.LoadEnv())
	require.NoError(t, err)
	var room models.Room
	require.NoError(t, db.First(&room).Error)

	bookingBase := time.Now().UnixNano()
	checkIn := time.Now().UTC().AddDate(2, 0, 0).Truncate(24 * time.Hour)
	checkOut := checkIn.AddDate(0, 0, 2)
	repository := NewRoomReservationRepository(db)
	results := make(chan error, 2)
	var waitGroup sync.WaitGroup
	for i := 0; i < 2; i++ {
		waitGroup.Add(1)
		go func(bookingID int64) {
			defer waitGroup.Done()
			reservation := &models.RoomReservation{RoomID: room.ID, BookingID: bookingID, CheckInDate: checkIn, CheckOutDate: checkOut, Status: models.RoomReservationStatusActive}
			_, _, createErr := repository.CreateReservation(reservation)
			results <- createErr
		}(bookingBase + int64(i))
	}
	waitGroup.Wait()
	close(results)

	successes := 0
	conflicts := 0
	for createErr := range results {
		if createErr == nil {
			successes++
		} else if createErr == ErrRoomReservationConflict {
			conflicts++
		}
	}
	require.Equal(t, 1, successes)
	require.Equal(t, 1, conflicts)
	db.Exec("DELETE FROM room_reservations WHERE booking_id IN (?, ?)", bookingBase, bookingBase+1)
}

func TestRoomReservationReleaseMakesRoomAvailable(t *testing.T) {
	if os.Getenv("HOTELS_RUN_INTEGRATION_TESTS") != "1" {
		t.Skip("set HOTELS_RUN_INTEGRATION_TESTS=1 to run against CockroachDB")
	}

	db, err := database.Init(config.LoadEnv())
	require.NoError(t, err)
	var room models.Room
	require.NoError(t, db.Where("active = ?", true).First(&room).Error)

	bookingID := time.Now().UnixNano()
	checkIn := time.Now().UTC().AddDate(3, 0, 0).Truncate(24 * time.Hour)
	checkOut := checkIn.AddDate(0, 0, 2)
	reservationRepository := NewRoomReservationRepository(db)
	reservation, _, err := reservationRepository.CreateReservation(&models.RoomReservation{RoomID: room.ID, BookingID: bookingID, CheckInDate: checkIn, CheckOutDate: checkOut, Status: models.RoomReservationStatusActive})
	require.NoError(t, err)
	defer db.Delete(&models.RoomReservation{}, "id = ?", reservation.ID)

	roomRepository := NewRoomRepository(db)
	available, err := roomRepository.GetAvailableRooms(room.HotelID, checkIn, checkOut, 1)
	require.NoError(t, err)
	require.NotContains(t, roomIDs(available), room.ID)

	released, err := reservationRepository.ReleaseByBookingID(bookingID)
	require.NoError(t, err)
	require.Equal(t, models.RoomReservationStatusReleased, released.Status)

	available, err = roomRepository.GetAvailableRooms(room.HotelID, checkIn, checkOut, 1)
	require.NoError(t, err)
	require.Contains(t, roomIDs(available), room.ID)
}

func roomIDs(rooms []models.Room) []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(rooms))
	for _, room := range rooms {
		ids = append(ids, room.ID)
	}
	return ids
}
