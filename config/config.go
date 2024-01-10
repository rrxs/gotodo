package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func SetupConfig() error {
	var err error

	db, err = setupSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetLogger(l string) *Logger {
	logger = NewLogger(l)
	return logger
}
