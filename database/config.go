// Package database
package database

import (
	"log"

	"mobileHost/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("databse/store.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error, could not connect to database")
	}

	// auto migrate project
	err = database.AutoMigrate((&models.Project{}))
	if err != nil {
		log.Fatal("Error, could not migrate database")
	}

	DB = database
}
