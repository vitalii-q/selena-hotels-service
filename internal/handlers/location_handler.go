package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitali-q/selena-hotels-service/internal/services"
)

type LocationHandler struct {
	service *services.LocationService
}

func NewLocationHandler(service *services.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func RegisterLocationRoutes(r *gin.RouterGroup, h *LocationHandler) {
	r.GET("/locations", h.GetCountriesWithCities)
}

func (h *LocationHandler) GetCountriesWithCities(c *gin.Context) {
	data, err := h.service.GetCountriesWithCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load locations",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
