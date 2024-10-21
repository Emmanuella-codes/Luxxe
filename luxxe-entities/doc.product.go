package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID            primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	Name          string              `json:"name" bson:"name"`
	Description   string              `json:"description" bson:"description"`
	Price         float64             `json:"price" bson:"price"`
	Category      string              `json:"category" bson:"category"`
	Quantity      int                 `json:"quantity" bson:"quantity"`
	CreatedAt     time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt" bson:"updatedAt"`
}

var ProductModel *Model[Product]
var ProductCollection *mongo.Collection

func initProduct() {
  ProductCollection = config.GetCollection(string(ModelNamesProduct))
  ProductModel = InitModel[Product](ModelNamesProduct, ProductCollection)
}
