package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // import registration for side effects
	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/handlers"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/sirupsen/logrus"
)

var DB *gorm.DB

func main() {
	// Create a new Gin instance
	r := gin.Default()

	//logrus.SetLevel(logrus.DebugLevel)           // Setting the logging level
	//logrus.SetFormatter(&logrus.TextFormatter{ 
	//	FullTimestamp: true,                       // Beautiful log output
	//})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, hotels-service!")
	})

	// Инициализация базы данных через GORM
	log.Println("Initializing DB...")
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("DB initialized successfully")

	//logrus.Debug("qwer1")

	// Checking the connection to the database
	r.GET("/health/db", func(c *gin.Context) {
		if database.DB == nil {
			log.Println("database.DB is nil")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB is nil"})
			return
    	}

		sqlDB, err := database.DB.DB() // get *sql.DB from GORM
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

		log.Println("DB ping successful")
		c.String(http.StatusOK, "Hotels-service: database connection OK ✅")
	})

	//logrus.Error("ests")
	//logrus.Debug("sfds")
	//logrus.Debug("Hotel service started")

	r.GET("/api/v1/locations", handlers.GetCountriesWithCities)

	handlers.RegisterHotelRoutes(r)

	// Configure the server to listen on port 8080
	if err := r.Run(":9064"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
