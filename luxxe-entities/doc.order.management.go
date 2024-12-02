package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderStatus string

const (
	OrderStatusPending 	 OrderStatus = "pending"
	OrderStatusShipped 	 OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
)

type OrderManagement struct {
	ID        			primitive.ObjectID 	`json:"_id" bson:"_id"`
	UserID    			primitive.ObjectID 	`json:"userID" bson:"userID"`
	CartID 					primitive.ObjectID  `json:"cartID" bson:"cart"`
	ShippingAddress string 							`json:"shippingAddress" bson:"shippingAddress"`
	PhoneNumber 		string 							`json:"phoneNumber" bson:"phoneNumber"`
	OrderStatus 		OrderStatus 				`json:"orderStatus" bson:"orderStatus"`
	CartTotal 			float64 						`json:"cartTotal" bson:"cartTotal"`
	CreatedAt 			time.Time          	`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			time.Time          	`json:"updatedAt" bson:"updatedAt"`
}

var OrderManagementModel *Model[OrderManagement]
var OrderManagementCollection *mongo.Collection

func initOrderManagement() {
	OrderManagementCollection = config.GetCollection(string(ModelNamesOrderManagement))
	OrderManagementModel = InitModel[OrderManagement](ModelNamesOrderManagement, OrderManagementCollection)
}
