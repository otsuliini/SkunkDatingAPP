package database

import (
	
	"log"

	"gorm.io/driver/sqlite" // SQLite driver
	"gorm.io/gorm"          // GORM ORM

	"github.com/yourusername/yourproject/internal/account" // Update this to the correct path
)

var DB *gorm.DB
var err error 
func InitDB() {
	DB, err = gorm.Open(sqlite.Open("raccoon.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

func updateDB() {
	DB.AutoMigrate(&account.Raccoon{}) // Update this to the correct path
}

func clearDB() {
	DB.Migrator().DropTable(&account.Raccoon{}) // Update this to the correct path
}

func getDB() *gorm.DB {
	return DB
}
