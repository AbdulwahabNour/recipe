package main

import (
	"log"

	"github.com/AbdulwahabNour/recipe/internal/server"
	"github.com/AbdulwahabNour/recipe/pkg/db/mongo"
)

func main() {

	mongoConn, err := mongo.NewMongoClient()
	if err != nil {
		panic(err)
	}
	mongoDB := mongoConn.Client.Database("recipes")

	server := server.NewServer(mongoDB)
	err = server.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
