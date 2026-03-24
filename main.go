package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // import registration for side effects
	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/handlers"
	"github.com/vitali-q/hotels-service/internal/services"
	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
	//"github.com/sirupsen/logrus"
)

func main() {
	// --- Router logs settings ---
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


	//logrus.SetLevel(logrus.DebugLevel)           // Setting the logging level
	//logrus.SetFormatter(&logrus.TextFormatter{ 
	//	FullTimestamp: true,                       // Beautiful log output
	//})

	// --- Router routes settings ---
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, hotels-service!")
	})

	// --- Database initialization ---
	log.Println("Initializing DB...")
	db, err := database.Init(); if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("DB initialized successfully")

	//logrus.Debug("qwer1")

	r.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "OK") })

	// Checking the connection to the database
	r.GET("/ready", func(c *gin.Context) {DBcheck(c, db)})

	//logrus.Error("ests")
	//logrus.Debug("sfds")
	//logrus.Debug("Hotel service started")

	hotelService := services.NewHotelService(db)
	hotelHandler := handlers.NewHotelHandler(hotelService)

	locationService := services.NewLocationService(db)
	locationHandler := handlers.NewLocationHandler(locationService)

	// API routers
	api := r.Group("/api/v1")
	handlers.RegisterHotelRoutes(api, hotelHandler)
	handlers.RegisterLocationRoutes(api, locationHandler)

	// Configure the server to listen on port 8080
	if err := r.Run(":9064"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func DBcheck(c *gin.Context, db *gorm.DB) {
		if db == nil {
			log.Println("database.DB is nil")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB is nil"})
			return
    	}

		sqlDB, err := db.DB() // get *sql.DB from GORM
		if err != nil {
			log.Printf("Failed to get raw DB: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get raw DB"})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			log.Printf("DB ping failed: %v\n", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database ping failed"})
			return
		}

		//log.Println("DB ping successful")
		c.String(http.StatusOK, "Hotels-service: database connection OK ✅")
	}