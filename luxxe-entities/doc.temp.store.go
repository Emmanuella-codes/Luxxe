package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TempStore struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Key            string             `json:"key" bson:"key"`
	Value          string             `json:"value" bson:"value"`
	ExpirationTime int                `json:"expirationTime" bson:"expirationTime"`
	BeginTime      int64              `json:"beginTime" bson:"beginTime"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var TempStoreModel *Model[TempStore]
var TempStoreCollection *mongo.Collection

func initTempStore() {
	TempStoreCollection = config.GetCollection(string(ModelNamesTempStore))
	TempStoreModel = InitModel[TempStore](ModelNamesTempStore, TempStoreCollection)
}
