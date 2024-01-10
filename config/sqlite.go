package config

import (
	"os"

	"github.com/rrxs/gotodo/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("creating database file")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(types.Todo{})
	if err != nil {
		logger.Errorf("auto migrate error: %v", err)
		return nil, err
	}

	return db, nil
}
