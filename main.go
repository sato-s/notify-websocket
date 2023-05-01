package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
	}
	NewServer(os.Getenv("PORT"))
}
