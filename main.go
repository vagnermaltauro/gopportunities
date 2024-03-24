package main

import (
	"github.com/vagnermaltauro/gopportunities/config"
	"github.com/vagnermaltauro/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization failed: %v", err)
		return
	}

	router.Initialize()
}
