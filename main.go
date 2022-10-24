package main

import (
	"log"

	"finall/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
