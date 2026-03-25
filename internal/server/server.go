package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NewHTTPServer creates configured HTTP server
func NewHTTPServer(port string, handler *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second, // protect from slowloris
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}