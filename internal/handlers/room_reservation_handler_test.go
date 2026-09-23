package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"github.com/vitali-q/selena-hotels-service/internal/services"
)

type reservationServiceStub struct {
	reservation *models.RoomReservation
	existed     bool
	err         error
}

func (s *reservationServiceStub) CreateReservation(roomID uuid.UUID, bookingID int64, checkInDate, checkOutDate time.Time, expiresAt *time.Time) (*models.RoomReservation, bool, error) {
	if s.err != nil {
		return nil, false, s.err
	}
	return s.reservation, s.existed, nil
}

func TestRoomReservationHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	reservationID := uuid.Must(uuid.NewV4())
	service := &reservationServiceStub{reservation: &models.RoomReservation{ID: reservationID, Status: models.RoomReservationStatusActive}}
	r := gin.New()
	RegisterRoomReservationRoutes(r, NewRoomReservationHandler(service))

	roomID := uuid.Must(uuid.NewV4())
	request := httptest.NewRequest("POST", "/internal/rooms/"+roomID.String()+"/reservations", strings.NewReader(`{"bookingId":42,"checkInDate":"2026-10-01","checkOutDate":"2026-10-03","expiresAt":"2026-09-30T12:00:00Z"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	require.Equal(t, 201, response.Code)
	require.Contains(t, response.Body.String(), reservationID.String())
}

func TestRoomReservationHandlerConflict(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := &reservationServiceStub{err: services.ErrReservationNotAvailable}
	r := gin.New()
	RegisterRoomReservationRoutes(r, NewRoomReservationHandler(service))

	roomID := uuid.Must(uuid.NewV4())
	request := httptest.NewRequest("POST", "/internal/rooms/"+roomID.String()+"/reservations", strings.NewReader(`{"bookingId":42,"checkInDate":"2026-10-01","checkOutDate":"2026-10-03","expiresAt":"2026-09-30T12:00:00Z"}`))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	require.Equal(t, 409, response.Code)
	require.Contains(t, response.Body.String(), "ROOM_NOT_AVAILABLE")
}
