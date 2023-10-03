package main

import (
	"log"

	"github.com/AbdulwahabNour/recipe/internal/server"
)

func main() {
	server := server.NewServer()
	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
