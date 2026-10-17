package services

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type ReservationExpiryRepository interface {
	ExpireReservations(time.Time) (int64, error)
}

var expiredReservationsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "hotels_room_reservations_expired_total",
	Help: "Total number of room reservations transitioned to EXPIRED.",
})

var registerExpiredMetric sync.Once

type ReservationExpiryWorker struct {
	repository ReservationExpiryRepository
	interval   time.Duration
}

func NewReservationExpiryWorker(repository ReservationExpiryRepository, interval time.Duration) *ReservationExpiryWorker {
	registerExpiredMetric.Do(func() { prometheus.MustRegister(expiredReservationsTotal) })
	return &ReservationExpiryWorker{repository: repository, interval: interval}
}

func (w *ReservationExpiryWorker) Run(ctx context.Context) {
	if w.interval <= 0 {
		return
	}
	w.expire(time.Now())
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			w.expire(now)
		}
	}
}

func (w *ReservationExpiryWorker) expire(now time.Time) {
	count, err := w.repository.ExpireReservations(now)
	if err != nil {
		log.Printf("failed to expire room reservations: %v", err)
		return
	}
	if count > 0 {
		expiredReservationsTotal.Add(float64(count))
	}
}

func ExpiredReservationsTotal() prometheus.Counter { return expiredReservationsTotal }
