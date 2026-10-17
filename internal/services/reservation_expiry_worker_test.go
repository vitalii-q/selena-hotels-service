package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vitali-q/selena-hotels-service/internal/models"
)

type expiryRepositoryStub struct {
	reservations []models.RoomReservation
	called       bool
}

func (s *expiryRepositoryStub) ExpireReservations(now time.Time) (int64, error) {
	s.called = true
	var count int64
	for i := range s.reservations {
		reservation := &s.reservations[i]
		if reservation.Status == models.RoomReservationStatusActive && reservation.ExpiresAt != nil && !reservation.ExpiresAt.After(now) {
			reservation.Status = models.RoomReservationStatusExpired
			count++
		}
	}
	return count, nil
}

func TestReservationExpiryWorkerExpiresReservations(t *testing.T) {
	expiresAt := time.Now().Add(-time.Minute)
	repository := &expiryRepositoryStub{reservations: []models.RoomReservation{
		{Status: models.RoomReservationStatusActive, ExpiresAt: &expiresAt},
		{Status: models.RoomReservationStatusActive, ExpiresAt: func() *time.Time { value := time.Now().Add(time.Hour); return &value }()},
	}}
	worker := NewReservationExpiryWorker(repository, time.Hour)
	worker.expire(time.Now())
	require.True(t, repository.called)
	require.Equal(t, models.RoomReservationStatusExpired, repository.reservations[0].Status)
	require.Equal(t, models.RoomReservationStatusActive, repository.reservations[1].Status)
}
