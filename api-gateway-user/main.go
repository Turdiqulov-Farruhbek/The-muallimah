package main

import (
	"muallimah-gateway/internal/app"
	"muallimah-gateway/internal/pkg/config"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}
