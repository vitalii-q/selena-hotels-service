package models

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func TestRoomReservationValidate(t *testing.T) {
	start := time.Date(2026, 9, 10, 0, 0, 0, 0, time.UTC)
	reservation := RoomReservation{
		RoomID: uuid.Must(uuid.NewV4()), BookingID: 42,
		CheckInDate: start, CheckOutDate: start.AddDate(0, 0, 2),
		Status: RoomReservationStatusActive,
	}

	require.NoError(t, reservation.Validate())
	reservation.CheckOutDate = start
	require.ErrorIs(t, reservation.Validate(), ErrRoomReservationDatesInvalid)
	reservation.CheckOutDate = start.AddDate(0, 0, 2)
	reservation.Status = "UNKNOWN"
	require.ErrorIs(t, reservation.Validate(), ErrRoomReservationStatusInvalid)
}

func TestRoomReservationOverlaps(t *testing.T) {
	start := time.Date(2026, 9, 10, 0, 0, 0, 0, time.UTC)
	reservation := RoomReservation{CheckInDate: start, CheckOutDate: start.AddDate(0, 0, 3), Status: RoomReservationStatusActive}

	require.True(t, reservation.Overlaps(start.AddDate(0, 0, 1), start.AddDate(0, 0, 2)))
	require.False(t, reservation.Overlaps(start.AddDate(0, 0, 3), start.AddDate(0, 0, 5)))
	reservation.Status = RoomReservationStatusExpired
	require.False(t, reservation.Overlaps(start.AddDate(0, 0, 1), start.AddDate(0, 0, 2)))
}
