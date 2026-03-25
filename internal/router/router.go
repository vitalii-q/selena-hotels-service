package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitali-q/hotels-service/internal/bootstrap"
	"github.com/vitali-q/hotels-service/internal/handlers"
	"github.com/vitali-q/hotels-service/internal/server/middleware"
)

func SetupRouter(deps *bootstrap.Dependencies) *gin.Engine {
	// --- Router initialization ---
	r := gin.New() // creating a router without the standard logger
	r.SetTrustedProxies(nil) // secure proxy configuration

	// --- Middleware ---
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// --- Root ---
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, hotels-service!")
	})

	// --- Health checks ---
	r.GET("/health", HealthHandler)
	r.GET("/ready", ReadyHandler(deps.DB))

	// --- API routes ---
	api := r.Group("/api/v1")

	handlers.RegisterHotelRoutes(api, deps.HotelHandler)
	handlers.RegisterLocationRoutes(api, deps.LocationHandler)

	return r
}