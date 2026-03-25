package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Recovery handles panics and prevents server crash
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {

		log.Printf("panic recovered: %v", recovered)

		c.AbortWithStatus(500)
	})
}