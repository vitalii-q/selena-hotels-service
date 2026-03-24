package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitali-q/hotels-service/internal/bootstrap"
	"github.com/vitali-q/hotels-service/internal/handlers"
)

func SetupRouter(deps *bootstrap.Dependencies) *gin.Engine {
	// --- Logs settings ---
	r := gin.New() // creating a router without the standard logger
	r.SetTrustedProxies(nil) // secure proxy configuration
	r.Use(gin.Recovery())    // Recover from panics to prevent server crash
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    gin.DefaultWriter,
		SkipPaths: []string{
			"/health", // health endpoint
			"/ready",    // enabling standard logs at this address
		},
	}))

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