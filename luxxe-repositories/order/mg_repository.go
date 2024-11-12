package order

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type mgRepository struct {
	log *log.Logger
}

func newMgRepository(log *log.Logger) OrderRepository {
	return &mgRepository{
		log: log,
	}
}

func (r *mgRepository) Create(ctx context.Context, order *entities.OrderManagement) (*entities.OrderManagement, error) {
	userID, shippingAddress, phoneNumber := order.UserID, order.ShippingAddress, order.PhoneNumber

	return entities.OrderManagementModel.Create(
		ctx,
		&bson.M{
			"userID": 				 userID,
			"shippingAddress": shippingAddress,
			"phoneNumber": 		 phoneNumber,
			"createdAt": 			 time.Now(),
		},
	)
}

func (r *mgRepository) GetOrder(ctx context.Context, userID string) (*entities.OrderManagement, int64, error) {
	userIDObj, _ := primitive.ObjectIDFromHex(userID)

	filter := &bson.M{"userID": userIDObj}

	orderCount, err := entities.OrderManagementCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var order entities.OrderManagement
	err = entities.OrderManagementCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, orderCount, nil
		}
		return nil, 0, err
	}

	return &order, orderCount, nil
}

func (r *mgRepository) QueryByUserID(ctx context.Context, userID string) (*entities.OrderManagement, error) {
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	filter := &primitive.M{"userID": userIDObj}
	order, err := entities.OrderManagementModel.FindOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find cart: %w", err)
	}

	return order, nil
}

func (r *mgRepository) CancelOrder(ctx context.Context, userID string) error {
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid format: %w", err)
	}

	_, err = entities.OrderManagementCollection.DeleteOne(ctx, bson.M{"userID": userIDObj})
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}

	return nil
}
