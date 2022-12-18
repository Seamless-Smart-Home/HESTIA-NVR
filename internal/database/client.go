package database

import (
	"log"

	"HESTIA/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("hestia.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	Instance = db

	log.Println("Connected to Database!")
	return nil
}

func Migrate() {
	Instance.AutoMigrate(
		&models.People{},
		&models.Areas{},
		&models.Cameras{},
		&models.Matches{},
	)
	log.Println("Database Migration Completed!")
}
