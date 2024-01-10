package handler

import (
	"github.com/rrxs/gotodo/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func SetupHandlers() {
	logger = config.GetLogger("handler")
	db = config.GetDB()

}
