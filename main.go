package main

import (
	"github.com/rrxs/gotodo/config"
	"github.com/rrxs/gotodo/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.SetupConfig()
	if err != nil {
		logger.Errorf("error setting up config: %v", err)
		return
	}

	router.Init()
}
