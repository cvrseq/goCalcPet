package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dataSrcName := "host=localhost user=postgres password=yourpassword dbname=postgres port=5000 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dataSrcName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return db, nil
}
