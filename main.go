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
	router.Init()
}
