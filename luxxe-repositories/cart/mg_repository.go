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

	cart_service "github.com/Emmanuella-codes/Luxxe/luxxe-cart/service"
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

const maxCartItems int = 30

func (r *mgRepository) AddToCart(ctx context.Context, userID string, productID string, quantity int, price float64) (*entities.Cart, error) {
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	productIDObj, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID: %w", err)
	}

	pipeline := []bson.M{
		{"$match": bson.M{"userID": userIDObj}},
		{"$project": bson.M{"itemCount": bson.M{"$size": "$items"}}},
	}
	cursor, err := entities.CartItemCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error counting items: %w", err)
	}
	defer cursor.Close(ctx)

	var cartInfo struct { ItemCount int}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&cartInfo); err != nil {
			return nil, fmt.Errorf("error decoding item count: %w", err)
		}
	}
	if cartInfo.ItemCount >= maxCartItems {
		return nil, fmt.Errorf("cart item limit of %d reached", maxCartItems)
	}

	itemTotalPrice := float64(quantity) * price

	filter := bson.M{
		"userID": userIDObj,
		"items.productID": productIDObj,
	}
	update := bson.M{
		"$inc": bson.M{"items.$.quantity": quantity},
		"$set": bson.M{
			"items.$.price": 			price,
			"items.$.totalPrice": itemTotalPrice,
			"updatedAt":     			time.Now(),
		},
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
		"$push": bson.M{
			"items": bson.M{ // Add new item to items array
				"productID": 	productIDObj,
				"quantity":  	quantity,
				"price": 		 	price,
				"totalPrice": itemTotalPrice,
		}},
	}

	opts.SetUpsert(true)
	err = entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		return nil, fmt.Errorf("failed to add new item to cart: %w", err)
	}

	totalAmount := cart_service.CalculateCartTotal(ctx, userIDObj)
	_, updateErr := entities.CartItemCollection.UpdateOne(ctx, bson.M{"userID": userIDObj}, bson.M{"$set": bson.M{"totalAmount": totalAmount}})
	if updateErr != nil {
		return nil, fmt.Errorf("error updating total amount: %w", updateErr)
	}
	updatedCart.TotalAmount = totalAmount

	return &updatedCart, nil
}

func (r *mgRepository) UpdateCartItem(ctx context.Context, userID string, productID string, quantity int, price float64) (*entities.Cart, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)
	productIDObj, _ := primitive.ObjectIDFromHex(productID)

	filter := bson.M{"userID": userIDObj, "items.productID": productIDObj}

	var existingCart entities.Cart
    err := entities.CartItemCollection.FindOne(ctx, filter).Decode(&existingCart)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, fmt.Errorf("item not found in the cart")
        }
        return nil, err
    }

		//get the current item's price and quantity
		var currentItem *entities.CartItem
    for _, item := range existingCart.Items {
        if item.ProductID == productIDObj {
            currentItem = &item
            break
        }
    }
    if currentItem == nil {
        return nil, fmt.Errorf("item not found in the cart")
    }

		oldAmount := currentItem.TotalPrice
    newAmount := float64(quantity) * price
    amountDifference := newAmount - oldAmount

	update := bson.M{
		"$set": bson.M{
			"items.$.quantity":   quantity,
			"items.$.price":      price,
			"items.$.totalPrice": newAmount,
			"updatedAt":          time.Now(),
		},
		"$inc": bson.M{
			"totalAmount": amountDifference, // Adjust the total amount
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	
	var updatedCart entities.Cart
	err = entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
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
	after := options.After             
	opts.ReturnDocument = &after

	var existingCart entities.Cart
	err := entities.CartItemCollection.FindOne(ctx, filter).Decode(&existingCart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("cart not found")
		}
		return nil, fmt.Errorf("error fetching cart: %w", err)
	}

	var removedItem *entities.CartItem
	for _, item := range existingCart.Items {
		if item.ProductID == productIDObj {
			removedItem = &item
			break
		}
	}
	if removedItem == nil {
		return nil, fmt.Errorf("item not found in the cart")
	}

	amountDifference := -removedItem.TotalPrice
	update["$inc"] = bson.M{"totalAmount": amountDifference}

	var updatedCart entities.Cart
	err = entities.CartItemCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCart)
	if err != nil {
		return nil, err
	}

	return &updatedCart, nil
}

func (r *mgRepository) GetCart(ctx context.Context, userID string) (*entities.Cart, int64, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)

	// skip, limit := misc.Pagination(misc.PaginationStruct{Page: page})

	filter := &bson.M{"userID": userIDObj}

	cartCount, err := entities.CartItemCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var cart entities.Cart
	err = entities.CartItemCollection.FindOne(ctx, filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, cartCount, nil
		}
		return nil, 0, err
	}
	
	return &cart, cartCount, nil
}

func (r *mgRepository) QueryByUserID(ctx context.Context, userID string) (*entities.Cart, error) {
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	
	filter := &primitive.M{"userID": userIDObj}
	cart, err := entities.CartItemModel.FindOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find cart: %w", err)
	}

	return cart, nil
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
