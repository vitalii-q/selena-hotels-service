package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"

	_ "github.com/lib/pq" // import registration for side effects
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitali-q/selena-hotels-service/internal/bootstrap"
	"github.com/vitali-q/selena-hotels-service/internal/router"
	"github.com/vitali-q/selena-hotels-service/internal/server"
)

func main() {
	// --- Bootstrap all dependencies ---
	deps, err := bootstrap.Init()
	if err != nil {
		log.Fatalf("Failed to bootstrap app: %v", err)
	}

	// --- Setup router ---
	r := router.SetupRouter(deps)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	if deps.Env.ReservationExpiryEnabled {
		workerContext, cancelWorker := context.WithCancel(context.Background())
		defer cancelWorker()
		go deps.ReservationExpiryWorker.Run(workerContext)
	}

	// --- Create HTTP server ---
	srv := server.NewHTTPServer(deps.Env.Port, r)

	// --- Run server with graceful shutdown ---
	server.Run(srv)
}
