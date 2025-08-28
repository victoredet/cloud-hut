// Package database
package database

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"mobileHost/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ensure the directory exists
	dbPath := "database/store.db"
	err := os.MkdirAll(filepath.Dir(dbPath), os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Error, could not connect to database:", err)
	}

	err = database.AutoMigrate(&models.Project{})
	if err != nil {
		log.Fatal("Error, could not migrate database:", err)
	}

	DB = database
}
