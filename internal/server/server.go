package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NewHTTPServer creates configured HTTP server with production timeouts
func NewHTTPServer(port string, handler *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second, // max time to read request (protection against slow POST/PUT)
		WriteTimeout: 10 * time.Second, // max time to write response (protection against hung connections)
		IdleTimeout:  60 * time.Second, // max time for keep-alive (keep-alive)
	}
}