package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HealthHandler returns simple health status
func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// ReadyHandler checks database connectivity
func ReadyHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		if db == nil {
			log.Println("database.DB is nil")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB is nil"})
			return
		}

		sqlDB, err := db.DB()
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

		c.String(http.StatusOK, "Hotels-service: database connection OK ✅")
	}
}