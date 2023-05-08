package main

import (
	"github.com/esceer/aoe-monk-it/internal/app"
	"github.com/esceer/aoe-monk-it/internal/config"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		panic(err)
	}
	app.RegisterHooks(cfg)
}
