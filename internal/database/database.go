package database

import (
	"github.com/dmitrie43/LibrarySearcherUser/internal/app/models"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func New() *DB {
	return &DB{Connection: NewPostgres()}
}

func Migrate(db *DB) {
	db.Connection.AutoMigrate(
		&models.Role{}, &models.User{},
	)
}
