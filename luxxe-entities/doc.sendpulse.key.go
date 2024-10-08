package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SendPulseKey struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	TokenType   string             `json:"token_type" bson:"token_type"`
	ExpiresIn   int                `json:"expires_in" bson:"expires_in"`
	AccessToken string             `json:"access_token" bson:"access_token"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var SendPulseKeyModel *Model[SendPulseKey]
var SendPulseKeyCollection *mongo.Collection

func initSendPulseKey() {
	SendPulseKeyCollection = config.GetCollection(string(ModelNamesSendPulseKey))
	SendPulseKeyModel = InitModel[SendPulseKey](ModelNamesSendPulseKey, SendPulseKeyCollection)
}
