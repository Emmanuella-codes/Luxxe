package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartItem struct {
	ProductID primitive.ObjectID `json:"productId" bson:"productId"`
	Quantity  int                `json:"quantity" bson:"quantity"`
}

type Cart struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	Items     []CartItem         `json:"items" bson:"items"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var CartItemModel *Model[AuditLog]
var CartItemCollection *mongo.Collection

func initCartItem() {
	AuditLogCollection = config.GetCollection(string(ModelNamesCart))
	AuditLogModel = InitModel[AuditLog](ModelNamesCart, CartItemCollection)
}
