package main

import (
	"backend/config"
	"backend/internal/app"
)

func main() {
	cfg, err := config.FillConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
