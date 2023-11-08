package database

import (
	"github.com/doganarif/todo/internal/models"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
