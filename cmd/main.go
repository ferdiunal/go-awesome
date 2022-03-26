package main

import (
	"log"

	"github.com/ferdiunal/go-awesome/cmd/src"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := src.NewAwesomeRepository()

	app.Run()
}
