package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-grpc-api/config"
	"go-grpc-api/database/migrations"
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

	migrations.Run(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
