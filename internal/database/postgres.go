package database

import (
	"log"

	"github.com/dmitrie43/LibrarySearcherUser/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	cnf := config.MustLoad()

	dsn := "user=" + cnf.DatabaseUser + " password=" + cnf.DatabasePassword + " dbname=" + cnf.DatabaseName + " port=" + cnf.DatabasePort
	connection, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatalf("Can't connect to database: %s", error)
	}

	return connection
}
