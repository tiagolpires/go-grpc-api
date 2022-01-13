package main

import (
	"go-grpc-api/database"
	"go-grpc-api/database/seeder/seeds"
	"log"
)

func main() {
	database.StartDB()
	db := database.GetDatabase()

	for _, seed := range seeds.Seeds {
		log.Printf("Running seed: %s", seed.Name)
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
