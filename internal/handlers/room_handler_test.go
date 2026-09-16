package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"github.com/vitali-q/selena-hotels-service/internal/services"
)

type roomServiceStub struct {
	room    *models.Room
	err     error
	deleted bool
}

func (s *roomServiceStub) CreateRoom(r *models.Room) (*models.Room, error) {
	if s.err != nil {
		return nil, s.err
	}
	s.room = r
	return r, nil
}
func (s *roomServiceStub) GetRoomByID(uuid.UUID) (*models.Room, error) { return s.room, s.err }
func (s *roomServiceStub) GetRoomsByHotelID(uuid.UUID) ([]models.Room, error) {
	if s.err != nil {
		return nil, s.err
	}
	return []models.Room{*s.room}, nil
}
func (s *roomServiceStub) UpdateRoom(r *models.Room) (*models.Room, error) { return r, s.err }
func (s *roomServiceStub) DeleteRoom(uuid.UUID) error                      { s.deleted = true; return s.err }
func (s *roomServiceStub) GetAvailableRooms(uuid.UUID, time.Time, time.Time, int) ([]models.Room, error) {
	if s.err != nil {
		return nil, s.err
	}
	if s.room == nil {
		return []models.Room{}, nil
	}
	return []models.Room{*s.room}, nil
}

func TestRoomHandlerCreateRoom(t *testing.T) {
	gin.SetMode(gin.TestMode)
	stub := &roomServiceStub{}
	r := gin.New()
	RegisterRoomRoutes(r.Group("/api/v1"), NewRoomHandler(stub))
	id := uuid.Must(uuid.NewV4())
	req := httptest.NewRequest("POST", "/api/v1/hotels/"+id.String()+"/rooms", strings.NewReader(`{"number":"101","type":"standard","capacity":2,"price_per_night":100}`))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 201, w.Code)
	require.Equal(t, "101", stub.room.Number)
	require.Equal(t, decimal.NewFromInt(100), stub.room.PricePerNight)
}

func TestRoomHandlerDeleteConflict(t *testing.T) {
	gin.SetMode(gin.TestMode)
	stub := &roomServiceStub{err: services.ErrActiveReservations}
	r := gin.New()
	RegisterRoomRoutes(r.Group("/api/v1"), NewRoomHandler(stub))
	id := uuid.Must(uuid.NewV4())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("DELETE", "/api/v1/rooms/"+id.String(), nil))
	require.Equal(t, 409, w.Code)
}

func TestRoomHandlerGetAvailableRooms(t *testing.T) {
	gin.SetMode(gin.TestMode)
	roomID := uuid.Must(uuid.NewV4())
	stub := &roomServiceStub{room: &models.Room{ID: roomID, HotelID: uuid.Must(uuid.NewV4()), Number: "101", Type: "standard", Capacity: 2, PricePerNight: decimal.NewFromInt(100), Active: true}}
	r := gin.New()
	RegisterRoomRoutes(r.Group("/api/v1"), NewRoomHandler(stub))
	w := httptest.NewRecorder()
	hotelID := uuid.Must(uuid.NewV4())
	request := httptest.NewRequest("GET", "/api/v1/rooms/available?hotelId="+hotelID.String()+"&checkInDate=2026-09-10&checkOutDate=2026-09-12&guests=2", nil)
	r.ServeHTTP(w, request)
	require.Equal(t, 200, w.Code)
}
