package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-grpc-api/config"
)

var db *gorm.DB

func StartDB() {
	database, err := gorm.Open(postgres.Open(config.DatabaseConnection), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	log.Println("Connected to Database 🚀")

}

func GetDatabase() *gorm.DB {
	return db
}
