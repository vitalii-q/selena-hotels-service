package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // import registration for side effects
	"github.com/vitali-q/hotels-service/internal/bootstrap"
	"github.com/vitali-q/hotels-service/internal/router"
	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
	//"github.com/sirupsen/logrus"
)

func main() {
	// --- Bootstrap all dependencies ---
	deps, err := bootstrap.Init()
	if err != nil {
		log.Fatalf("Failed to bootstrap app: %v", err)
	}

	// --- Setup router ---
	r := router.SetupRouter(deps)

	// --- Start server ---
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