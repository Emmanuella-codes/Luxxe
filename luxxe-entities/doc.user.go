package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID                   primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname            string             `json:"firstname" bson:"firstname"`
	Lastname             string             `json:"lastname" bson:"lastname"`
	Email                string             `json:"email" bson:"email"`
	EmailYetToBeVerified string             `json:"emailYetToBeVerified" bson:"emailYetToBeVerified"`
	Password             string             `json:"password" bson:"password"`
	CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var UserModel *Model[User]
var UserCollection *mongo.Collection

func initUser() {
	UserCollection = config.GetCollection(string(ModelNamesUser))
	UserModel = InitModel[User](ModelNamesUser, UserCollection)
}
