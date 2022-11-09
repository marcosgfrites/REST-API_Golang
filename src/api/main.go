package main

import (
	"github.com/marcosgfrites/REST-API_Golang/src/databases"
	"github.com/marcosgfrites/REST-API_Golang/src/handlers"
	"log"
)

func main() {
	if databases.ConnectionCheck() == 0 {
		log.Fatal("Database connection failed")
		return
	}

	handlers.Handlers()
}
