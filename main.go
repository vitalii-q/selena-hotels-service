package main

import (
	"log"

	_ "github.com/lib/pq" // import registration for side effects
	"github.com/vitali-q/hotels-service/internal/bootstrap"
	"github.com/vitali-q/hotels-service/internal/router"
	"github.com/vitali-q/hotels-service/internal/server"
)

func main() {
	// --- Bootstrap all dependencies ---
	deps, err := bootstrap.Init()
	if err != nil {
		log.Fatalf("Failed to bootstrap app: %v", err)
	}

	// --- Setup router ---
	r := router.SetupRouter(deps)

		// --- Create HTTP server ---
	srv := server.NewHTTPServer(deps.Env.Port, r)

	// --- Run server with graceful shutdown ---
	server.Run(srv)
}