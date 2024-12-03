package service

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateCartTotal(ctx context.Context, userID primitive.ObjectID) float64 {
	filter := bson.M{"userID": userID}
	var cart entities.Cart
	err := entities.CartItemCollection.FindOne(ctx, filter).Decode(&cart)
	if err != nil {
		return 0 // Handle error appropriately in production
	}

	totalAmount := 0.0
	for _, item := range cart.Items {
		totalAmount += item.TotalPrice
	}

	return totalAmount
}
