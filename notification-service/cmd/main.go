package main

import (
	"gitlab.com/muallimah/notification_service/internal/pkg/config"
	"gitlab.com/muallimah/notification_service/internal/app"
)

func main() {
	//configurations
	cfg := config.Load()

	// Run
	app.Run(&cfg)
}
