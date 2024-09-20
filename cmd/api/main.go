package main

import (
	"log"

	"github.com/rikughi/go-quick-start/internal/config"
)

func main() {
	viper := config.NewViper()
	app := config.NewApp(&config.App{
		Config: viper,
	})

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
