package main

import (
	"log"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/config"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
	app.Run(cfg)
}
