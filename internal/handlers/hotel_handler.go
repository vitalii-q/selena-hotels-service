package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/models"
	"github.com/vitali-q/hotels-service/internal/services"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	service *services.HotelService
}

func NewHotelHandler(service *services.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func RegisterHotelRoutes(r *gin.RouterGroup, h *HotelHandler) {
    hotels := r.Group("/hotels")
    {
        hotels.POST("", h.CreateHotel)
        hotels.GET("/:id", h.GetHotelByID)
        hotels.PUT("/:id", h.UpdateHotel)
        hotels.DELETE("/:id", h.DeleteHotel)

        hotels.GET("", h.GetHotels)
    }
}

func (h *HotelHandler) CreateHotel(c *gin.Context) {
    var hotel models.Hotel
    if err := c.ShouldBindJSON(&hotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdHotel, err := h.service.CreateHotel(&hotel)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdHotel)
}

func (h *HotelHandler) GetHotelByID(c *gin.Context) {
    uuidParam := c.Param("id")
    hotelID, err := uuid.FromString(uuidParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    hotel, err := h.service.GetHotelByID(hotelID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hotel)
}


func (h *HotelHandler) UpdateHotel(c *gin.Context) {
    uuidParam := c.Param("id")
    hotelID, err := uuid.FromString(uuidParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    var hotel models.Hotel
    if err := c.ShouldBindJSON(&hotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedHotel, err := h.service.UpdateHotel(hotelID, &hotel)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedHotel)
}


func (h *HotelHandler) DeleteHotel(c *gin.Context) {
    uuidParam := c.Param("id")
    hotelID, err := uuid.FromString(uuidParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    err = h.service.DeleteHotel(hotelID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

func (h *HotelHandler) GetHotels(c *gin.Context) {
	hotels, err := h.service.GetAllHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotels)
}
