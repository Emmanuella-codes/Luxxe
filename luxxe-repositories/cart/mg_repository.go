package cart

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-shared/misc"
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
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	productIDObj, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID: %w", err)
	}

	filter := bson.M{
		"userID": userIDObj,
		"items.productID": productIDObj,
	}
	update := bson.M{
		"$set": bson.M{"updatedAt": time.Now()},
		"$inc": bson.M{"items.$.quantity": quantity},
	}

	// Attempt to update an existing item
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedCart entities.Cart
	err = entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err == nil {
		// Item exists and quantity was incremented successfully, return updated cart
		return &updatedCart, nil
	}

	// Step 2: If no item found, add the product as a new item
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("error updating cart item: %w", err)
	}
		// Update the filter to find the user's cart (no product filter)
	filter = bson.M{"userID": userIDObj}
	update = bson.M{
		"$setOnInsert": bson.M{"createdAt": time.Now()},
		"$push": bson.M{"items": bson.M{ // Add new item to items array
			"productID": productIDObj,
			"quantity":  quantity,
		}},
	}

	opts.SetUpsert(true)
	err = entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		return nil, fmt.Errorf("failed to add new item to cart: %w", err)
	}

	return &updatedCart, nil
}

func (r *mgRepository) UpdateCartItem(ctx context.Context, userID string, productID string, quantity int) (*entities.Cart, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	productIDObj, _ := primitive.ObjectIDFromHex(productID)

	filter := bson.M{"userID": userIDObj, "items.productID": productIDObj}
	update := bson.M{
		"$inc": bson.M{"items.$.quantity": quantity,},
		"$set": bson.M{"updatedAt": time.Now()},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	
	var updatedCart entities.Cart
	err := entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("item not found in the cart")
		}
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

func (r *mgRepository) GetCart(ctx context.Context, userID string, page int) (*[]entities.Cart, int64, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)

	skip, limit := misc.Pagination(misc.PaginationStruct{Page: page})

	filter := &bson.M{"userID": userIDObj}

	cartCount, err := entities.CartItemCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := entities.CartItemCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var cart []entities.Cart = []entities.Cart{}
	if err = cursor.All(ctx, &cart); err != nil {
		return nil, 0, err
	}
	return &cart, cartCount, nil
}

func (r *mgRepository) ClearCart(ctx context.Context, userID string) error {
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid format: %w", err)
	}
	
	_, err = entities.CartItemCollection.DeleteOne(ctx, bson.M{"userID": userIDObj})
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}

	return nil
}
