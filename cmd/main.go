package main

import (
	"log"

	"github.com/rafliputraa/petstore/config"
	"github.com/rafliputraa/petstore/internal/app"
)

func main() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
