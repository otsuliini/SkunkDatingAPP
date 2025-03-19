package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/otsuliini/SkunkDatingAPP/internal/user"
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
	DB.AutoMigrate(&user.Raccoon{}) // Ensure Raccoon struct is defined in user package
}

func clearDB() {
	DB.Migrator().DropTable(&user.Raccoon{}) // Ensure Raccoon struct is defined in user package
}

func getDB() *gorm.DB {
	return DB
}
