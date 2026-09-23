package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"github.com/vitali-q/selena-hotels-service/internal/services"
	"gorm.io/gorm"
)

type RoomReservationService interface {
	CreateReservation(uuid.UUID, int64, time.Time, time.Time, *time.Time) (*models.RoomReservation, bool, error)
}

type RoomReservationHandler struct {
	service RoomReservationService
}

func NewRoomReservationHandler(service RoomReservationService) *RoomReservationHandler {
	return &RoomReservationHandler{service: service}
}

func RegisterRoomReservationRoutes(r *gin.Engine, h *RoomReservationHandler) {
	r.POST("/internal/rooms/:roomId/reservations", h.CreateReservation)
}

type createRoomReservationRequest struct {
	BookingID    int64  `json:"bookingId"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	ExpiresAt    string `json:"expiresAt"`
}

type roomReservationResponse struct {
	ReservationID uuid.UUID `json:"reservationId"`
	Status        string    `json:"status"`
}

func (h *RoomReservationHandler) CreateReservation(c *gin.Context) {
	roomID, err := uuid.FromString(c.Param("roomId"))
	if err != nil {
		roomError(c, http.StatusBadRequest, "INVALID_UUID", "Invalid roomId UUID")
		return
	}

	var request createRoomReservationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		roomError(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	checkInDate, err := time.Parse("2006-01-02", request.CheckInDate)
	if err != nil {
		roomError(c, http.StatusBadRequest, "INVALID_DATE", "checkInDate must use YYYY-MM-DD format")
		return
	}
	checkOutDate, err := time.Parse("2006-01-02", request.CheckOutDate)
	if err != nil {
		roomError(c, http.StatusBadRequest, "INVALID_DATE", "checkOutDate must use YYYY-MM-DD format")
		return
	}
	expiresAt, err := time.Parse(time.RFC3339, request.ExpiresAt)
	if err != nil {
		roomError(c, http.StatusBadRequest, "INVALID_DATE", "expiresAt must use RFC3339 format")
		return
	}

	reservation, existed, err := h.service.CreateReservation(roomID, request.BookingID, checkInDate, checkOutDate, &expiresAt)
	if err != nil {
		h.handleReservationError(c, err)
		return
	}
	status := http.StatusCreated
	if existed {
		status = http.StatusOK
	}
	c.JSON(status, roomReservationResponse{ReservationID: reservation.ID, Status: reservation.Status})
}

func (h *RoomReservationHandler) handleReservationError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrReservationNotAvailable):
		roomError(c, http.StatusConflict, "ROOM_NOT_AVAILABLE", err.Error())
	case errors.Is(err, gorm.ErrRecordNotFound):
		roomError(c, http.StatusNotFound, "ROOM_NOT_FOUND", "Room not found")
	case errors.Is(err, services.ErrReservationRoomRequired), errors.Is(err, services.ErrReservationBookingRequired), errors.Is(err, services.ErrReservationDatesInvalid):
		roomError(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
	default:
		roomError(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}
}
