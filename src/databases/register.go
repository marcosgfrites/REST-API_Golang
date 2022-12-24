package databases

import (
	"context"
	"fmt"
	"github.com/marcosgfrites/REST-API_Golang/src/models"
	"github.com/marcosgfrites/REST-API_Golang/src/utils/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

// UserSave inserts the user data into Database
func UserSave(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MDBConnection.Database(fmt.Sprintf("%s", os.Getenv("MONGODB_DATABASE")))
	collection := database.Collection("users")

	user.Password, _ = PasswordEncrypt(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", ok, fmt.Errorf(common.PrimitiveObjectIDError)
	}

	return objectID.String(), true, nil
}
