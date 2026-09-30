package repository

import (
	"os"
	"sync"
	"testing"
	"time"

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
