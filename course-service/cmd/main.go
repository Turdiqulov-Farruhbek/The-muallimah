package main

import (
	"gitlab.com/muallimah/course_service/internal/app"
	"gitlab.com/muallimah/course_service/internal/pkg/config"
)

func main() {
	config := config.Load()

	app.Run(&config)
}
