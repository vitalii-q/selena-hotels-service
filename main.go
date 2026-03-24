package main

import (
	"log"

	_ "github.com/lib/pq" // import registration for side effects
	"github.com/vitali-q/hotels-service/internal/bootstrap"
	"github.com/vitali-q/hotels-service/internal/router"
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
	if err := r.Run(":" + deps.Config.Port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}