package models

import (
	"fmt"
	"latihan-testing/models/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSLMODE")

	var err error
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbPass, dbUser, dbName, dbSSL,
	)

	entity.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cannot Connet Your DB")
	}
	return entity.DB
}

func CloseDatabaseConnection(db *gorm.DB) {
	DB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	DB.Close()
}
