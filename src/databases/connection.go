package databases

import (
	"context"
	"fmt"
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

	log.Println("The connection to the database is successful")
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

func getURLConnectionFromEnv() string {
	databaseURL := fmt.Sprintf("mongodb+srv://marcosdetti:%s@rest-api-golang-mgfd.do6n5si.mongodb.net/test", os.Getenv("MONGODB_PASSWORD"))
	return databaseURL
}
