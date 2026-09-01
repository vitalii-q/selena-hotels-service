package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/vitali-q/selena-hotels-service/internal/dto"
	"github.com/vitali-q/selena-hotels-service/internal/models"
	"github.com/vitali-q/selena-hotels-service/internal/services"
	"gorm.io/gorm"
)

type RoomService interface {
	CreateRoom(*models.Room) (*models.Room, error)
	GetRoomByID(uuid.UUID) (*models.Room, error)
	GetRoomsByHotelID(uuid.UUID) ([]models.Room, error)
	UpdateRoom(*models.Room) (*models.Room, error)
	DeleteRoom(uuid.UUID) error
}

type RoomHandler struct{ service RoomService }

func NewRoomHandler(service RoomService) *RoomHandler { return &RoomHandler{service: service} }

func RegisterRoomRoutes(r *gin.RouterGroup, h *RoomHandler) {
	r.POST("/hotels/:id/rooms", h.CreateRoom)
	r.GET("/hotels/:id/rooms", h.GetRoomsByHotel)
	r.GET("/rooms/:roomId", h.GetRoom)
	r.PATCH("/rooms/:roomId", h.UpdateRoom)
	r.DELETE("/rooms/:roomId", h.DeleteRoom)
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func roomError(c *gin.Context, status int, code, message string) {
	c.JSON(status, errorResponse{Error: code, Message: message})
}
func parseRoomUUID(c *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.FromString(c.Param(name))
	if err != nil {
		roomError(c, 400, "INVALID_UUID", "Invalid UUID")
		return uuid.Nil, false
	}
	return id, true
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	hotelID, ok := parseRoomUUID(c, "id")
	if !ok {
		return
	}
	var req dto.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		roomError(c, 400, "VALIDATION_ERROR", err.Error())
		return
	}
	room, err := h.service.CreateRoom(&models.Room{HotelID: hotelID, Number: strings.TrimSpace(req.Number), Type: strings.TrimSpace(req.Type), Capacity: req.Capacity, PricePerNight: req.PricePerNight, Active: true})
	if err != nil {
		h.handleRoomError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.MapRoom(room))
}

func (h *RoomHandler) GetRoomsByHotel(c *gin.Context) {
	hotelID, ok := parseRoomUUID(c, "id")
	if !ok {
		return
	}
	rooms, err := h.service.GetRoomsByHotelID(hotelID)
	if err != nil {
		h.handleRoomError(c, err)
		return
	}
	result := make([]dto.RoomResponse, 0, len(rooms))
	for i := range rooms {
		result = append(result, dto.MapRoom(&rooms[i]))
	}
	c.JSON(http.StatusOK, result)
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	id, ok := parseRoomUUID(c, "roomId")
	if !ok {
		return
	}
	room, err := h.service.GetRoomByID(id)
	if err != nil {
		h.handleRoomError(c, err)
		return
	}
	c.JSON(200, dto.MapRoom(room))
}

func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	id, ok := parseRoomUUID(c, "roomId")
	if !ok {
		return
	}
	room, err := h.service.GetRoomByID(id)
	if err != nil {
		h.handleRoomError(c, err)
		return
	}
	var req dto.UpdateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		roomError(c, 400, "VALIDATION_ERROR", err.Error())
		return
	}
	if req.Number != nil {
		room.Number = strings.TrimSpace(*req.Number)
	}
	if req.Type != nil {
		room.Type = strings.TrimSpace(*req.Type)
	}
	if req.Capacity != nil {
		room.Capacity = *req.Capacity
	}
	if req.PricePerNight != nil {
		room.PricePerNight = *req.PricePerNight
	}
	if req.Active != nil {
		room.Active = *req.Active
	}
	updated, err := h.service.UpdateRoom(room)
	if err != nil {
		h.handleRoomError(c, err)
		return
	}
	c.JSON(200, dto.MapRoom(updated))
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	id, ok := parseRoomUUID(c, "roomId")
	if !ok {
		return
	}
	if err := h.service.DeleteRoom(id); err != nil {
		h.handleRoomError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *RoomHandler) handleRoomError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrHotelNotFound):
		roomError(c, 404, "HOTEL_NOT_FOUND", err.Error())
	case errors.Is(err, services.ErrActiveReservations):
		roomError(c, 409, "ACTIVE_RESERVATIONS", err.Error())
	case errors.Is(err, gorm.ErrRecordNotFound), errors.Is(err, services.ErrRoomNotFound):
		roomError(c, 404, "ROOM_NOT_FOUND", "Room not found")
	case errors.Is(err, services.ErrRoomHotelRequired), errors.Is(err, services.ErrRoomNumberRequired), errors.Is(err, services.ErrRoomTypeRequired), errors.Is(err, services.ErrRoomCapacityInvalid), errors.Is(err, services.ErrRoomPricePerNightInvalid):
		roomError(c, 400, "VALIDATION_ERROR", err.Error())
	default:
		roomError(c, 500, "INTERNAL_ERROR", err.Error())
	}
}
