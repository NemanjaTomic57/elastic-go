// Package es initializes the connection to the elasticsearch docker container
package es

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
)

var Client *elasticsearch.Client

func Init() {
	var err error

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ES_LOCAL_URL"),
		},
		APIKey: os.Getenv("ES_LOCAL_API_KEY"),
	}

	Client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Init elasticsearch client: %v", err)
	}
}
