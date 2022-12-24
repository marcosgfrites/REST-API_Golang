package main

import (
	"github.com/marcosgfrites/REST-API_Golang/src/databases"
	"github.com/marcosgfrites/REST-API_Golang/src/handlers"
	"github.com/marcosgfrites/REST-API_Golang/src/utils/common"
	"log"
)

func main() {
	if databases.ConnectionCheck() == 0 {
		log.Fatal(common.DatabaseConnectionFailed)
		return
	}

	handlers.Handlers()
}
