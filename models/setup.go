package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	// "gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}

// gorm.io/driver/sqlite
