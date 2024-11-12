package order

import (
	"context"

	"github.com/go-kit/log"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type OrderRepository interface {
	Create(ctx context.Context,  order *entities.OrderManagement) (*entities.OrderManagement, error)
	GetOrder(ctx context.Context, userID string) (*entities.OrderManagement, int64, error)
	QueryByUserID(ctx context.Context, userID string) (*entities.OrderManagement, error)
	CancelOrder(ctx context.Context, userID string) error
}

var OrderRepo OrderRepository 

func InitOrderRepo(logger *log.Logger) {
	OrderRepo = newMgRepository(logger)
}
