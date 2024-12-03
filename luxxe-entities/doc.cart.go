package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartItem struct {
	ProductID 	primitive.ObjectID `json:"productID" bson:"productID"`
	Quantity  	int                `json:"quantity" bson:"quantity"`
	Price     	float64            `json:"price" bson:"price"`
	TotalPrice 	float64 					 `json:"totalPrice" bson:"totalPrice"`
}

type Cart struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserID      primitive.ObjectID `json:"userID" bson:"userID"`
	Items       []CartItem         `json:"items" bson:"items"`
	TotalAmount float64						 `json:"totalAmount" bson:"totalAmount"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var CartItemModel *Model[Cart]
var CartItemCollection *mongo.Collection

func initCartItem() {
	CartItemCollection = config.GetCollection(string(ModelNamesCart))
	CartItemModel = InitModel[Cart](ModelNamesCart, CartItemCollection)
}
