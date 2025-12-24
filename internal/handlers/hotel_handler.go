package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/vitali-q/hotels-service/internal/models"
	"github.com/vitali-q/hotels-service/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterHotelRoutes(r *gin.Engine) {
    hotels := r.Group("/hotels")
    {
        hotels.POST("", CreateHotel)
        hotels.GET("/:id", GetHotelByID)
        hotels.PUT("/:id", UpdateHotel)
        hotels.DELETE("/:id", DeleteHotel)

        hotels.GET("", GetHotels)
    }
}

func CreateHotel(c *gin.Context) {
    var hotel models.Hotel
    if err := c.ShouldBindJSON(&hotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdHotel, err := services.CreateHotel(&hotel)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdHotel)
}

func GetHotelByID(c *gin.Context) {
    uuidParam := c.Param("id")
    hotelID, err := uuid.FromString(uuidParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    hotel, err := services.GetHotelByID(hotelID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hotel)
}


func UpdateHotel(c *gin.Context) {
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

    updatedHotel, err := services.UpdateHotel(hotelID, &hotel)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedHotel)
}


func DeleteHotel(c *gin.Context) {
    uuidParam := c.Param("id")
    hotelID, err := uuid.FromString(uuidParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    err = services.DeleteHotel(hotelID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

func GetHotels(c *gin.Context) {
	hotels, err := services.GetAllHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotels)
}
