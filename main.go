package main

import (
	"log"

	"github.com/ferdiunal/go-awesome/src"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	miner := src.NewAwesomeRepository()

	miner.Run()
}