package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitali-q/hotels-service/internal/services"
)

func RegisterLocationRoutes(r *gin.RouterGroup) {
	r.GET("/locations", GetCountriesWithCities)
}

func GetCountriesWithCities(c *gin.Context) {
	data, err := services.GetCountriesWithCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load locations",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
