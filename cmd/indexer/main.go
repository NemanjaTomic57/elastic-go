package main

import (
	"log"

	"example.com/elastic-go/internal/db"
	"example.com/elastic-go/internal/es"
	"example.com/elastic-go/internal/indexer"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	es.Init()
	db.Init("elastic:password@tcp(127.0.0.1:3306)/main")

	indexer.IndexVerzeichnisse()

	log.Println("Done")
}
