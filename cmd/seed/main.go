package main

import (
	"log"

	"github.com/vitali-q/hotels-service/internal/database/seeds"
)

func main() {
	log.Println("🚀 Running hotels seeds...")
	seeds.RunSeeds()
}
