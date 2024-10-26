package cart

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	filter := bson.M{"userID": userIDObj, "items.productID": productIDObj}
	update := bson.M{
		"$set": bson.M{
			"items.$.quantity": quantity,
			"updatedAt": 				time.Now(),
		},
	}
	opts := options.FindOneAndUpdate()
	after := options.After             // ReturnDocument should be a pointer to ReturnDocument
	opts.ReturnDocument = &after

	var updatedCart entities.Cart
	err := entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		return nil, err
	}

	return &updatedCart, nil
}

func (r *mgRepository) RemoveFromCart(ctx context.Context, userID string, productID string) (*entities.Cart, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	productIDObj, _ := primitive.ObjectIDFromHex(productID)

	filter := bson.M{"userID": userIDObj}
	update := bson.M{
		"$pull": bson.M{
			"items": bson.M{"productID": productIDObj},
		},
		"$set": bson.M{"updatedAt": time.Now(),},
	}

	opts := options.FindOneAndUpdate()
	after := options.After             // ReturnDocument should be a pointer to ReturnDocument
	opts.ReturnDocument = &after

	var updatedCart entities.Cart
	err := entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		return nil, err
	}

	return &updatedCart, nil
}

func (r *mgRepository) GetCart(ctx context.Context, userID string) (*entities.Cart, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	
	var cart entities.Cart
	err := entities.CartItemCollection.FindOne(ctx, bson.M{"userID": userIDObj}).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newCart := &entities.Cart{
				ID: 			 primitive.NewObjectID(),
				UserID: 	 userIDObj,
				Items: 		 []entities.CartItem{},
				CreatedAt: time.Now(),
			}
			_, err := entities.CartItemCollection.InsertOne(ctx, newCart)
			if err != nil {
				return nil, err
			}
			return newCart, nil
		}
		return nil, err
	}

	return &cart, nil
}

func (r *mgRepository) ClearCart(ctx context.Context, userID string) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	
	entities.CartItemCollection.DeleteOne(ctx, bson.M{"userID": userIDObj})
}
