package handlers

import (
	
    "github.com/vitali-q/hotels-service/internal/models"
    "github.com/vitali-q/hotels-service/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func RegisterHotelRoutes(r *gin.Engine) {
    hotels := r.Group("/hotels")
    {
        hotels.POST("", CreateHotel)
        hotels.GET("", GetHotels)
        hotels.GET("/:id", GetHotelByID)
        hotels.PUT("/:id", UpdateHotel)
        hotels.DELETE("/:id", DeleteHotel)
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

func GetHotels(c *gin.Context) {
    hotels, err := services.GetAllHotels()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hotels)
}

func GetHotelByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    hotel, err := services.GetHotelByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hotel)
}

func UpdateHotel(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var hotel models.Hotel

    if err := c.ShouldBindJSON(&hotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedHotel, err := services.UpdateHotel(uint(id), &hotel)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedHotel)
}

func DeleteHotel(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := services.DeleteHotel(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}
