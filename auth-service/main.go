package main

import (
	"auth-service/internal/app"
	"auth-service/internal/pkg/config"
)

func main() {
	cf := config.Load()
	app.Run(cf)
}
