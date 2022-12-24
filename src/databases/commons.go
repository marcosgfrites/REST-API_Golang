package databases

import (
	"context"
	"fmt"
	"github.com/marcosgfrites/REST-API_Golang/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

// UserAlreadyExists validates if already exist another registered user with the same email
func UserAlreadyExists(userEmail string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MDBConnection.Database(fmt.Sprintf("%s", os.Getenv("MONGODB_DATABASE")))
	collection := database.Collection("users")

	condition := bson.M{"email": userEmail}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	id := result.Id.Hex()
	if err != nil {
		return result, false, id
	}

	return result, true, id
}
