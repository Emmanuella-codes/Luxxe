package cart

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type mgRepository struct {
	log *log.Logger
}

func newMgRepository(log *log.Logger) CartRepository {
	return &mgRepository{
		log: log,
	}
}

func (r *mgRepository) AddToCart(ctx context.Context, userID string, productID string, quantity int) (*entities.Cart, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	productIDObj, _ := primitive.ObjectIDFromHex(productID)

	filter := bson.M{
		"userID": userIDObj,
	}
	update := bson.M{
		"$setOnInsert": bson.M{
			"createdAt": time.Now(),
		},
		"$push": bson.M{
			"items": bson.M{
				"productID": productIDObj,
				"quantity": quantity,
			},
		},
	}

	opts := options.FindOneAndUpdate()
	upsert := true                     	// Upsert should be a pointer to a bool
	opts.Upsert = &upsert              	// Pass the pointer to Upsert

	after := options.After             // ReturnDocument should be a pointer to ReturnDocument
	opts.ReturnDocument = &after			// Pass the pointer to ReturnDocument

	var updatedCart entities.Cart
	err := entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		return nil, err
	}

	return &updatedCart, nil
}

func (r *mgRepository) UpdateCartItem(ctx context.Context, userID string, productID string, quantity int) (*entities.Cart, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	productIDObj, _ := primitive.ObjectIDFromHex(productID)

	
}

func (r *mgRepository) RemoveFromCart(ctx context.Context, userID string, productID string) (*entities.Cart, error) {
	
}

func (r *mgRepository) GetCart(ctx context.Context, userID string) (*entities.Cart, error) {
	
}

func (r *mgRepository) ClearCart(ctx context.Context, userID string) error {
	
}
