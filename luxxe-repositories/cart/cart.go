package cart

import (
	"context"

	"github.com/go-kit/log"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userID string, productID string, quantity int, price float64) (*entities.Cart, error)
	UpdateCartItem(ctx context.Context, userID string, productID string, quantity int, price float64) (*entities.Cart, error)
	RemoveFromCart(ctx context.Context, userID string, productID string) (*entities.Cart, error)
	QueryByUserID(ctx context.Context, userID string) (*entities.Cart, error)
	GetCart(ctx context.Context, userID string) (*entities.Cart, int64, error)
	ClearCart(ctx context.Context, userID string) error
}

var CartRepo CartRepository

func InitCartRepo(logger *log.Logger) {
	CartRepo = newMgRepository(logger)
}
