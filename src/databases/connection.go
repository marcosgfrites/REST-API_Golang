package databases

import (
	"context"
	"fmt"
	"github.com/marcosgfrites/REST-API_Golang/src/utils/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// MDBConnection contains the mongo connection client
var MDBConnection = ConnectToDB()

var clientOptions = options.Client().ApplyURI(getURLConnectionFromEnv())

// ConnectToDB allows to connect to the database
func ConnectToDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println(common.DatabaseConnectionSuccessful)
	return client
}

// ConnectionCheck do ping to the database
func ConnectionCheck() int {
	err := MDBConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}

// getURLConnectionFromEnv is a private function to get the URL connection by environment variable
func getURLConnectionFromEnv() string {
	databaseURL := fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("MONGODB_USER"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_API"))
	return databaseURL
}
