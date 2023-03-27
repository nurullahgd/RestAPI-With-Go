package config

import (
	"fmt"
	"time"

	"github.com/nurullahgd/RestAPI-With-Go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=PostgreGoRestAPI port=9123 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	start := time.Now()
	for sqlDB.Ping() != nil {
		if time.Now().After(start.Add(10 * time.Second)) {
			return nil, fmt.Errorf("Failed to connect DB after 10 seconds")
		}
	}
	fmt.Println("Connected to DB")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}
