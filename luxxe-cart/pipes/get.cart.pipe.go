package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type CartItems struct {
	CartItemCount int64						 `json:"cartItemCount"`
	CartItems     *[]entities.Cart `json:"cartItems"`
}

func GetCartPipe(ctx context.Context, dto *dtos.GetCartDTO) *shared.PipeRes[CartItems] {
	userID, page := dto.UserID, dto.Page

	cartItems, cartItemCount, _ := cart_repo.CartRepo.GetCart(ctx, userID, page)

	return &shared.PipeRes[CartItems]{
		Success: true,
		Message: messages.SUCCESS_GET_CART_ITEMS,
		Data:    &CartItems{
			CartItemCount: cartItemCount,
			CartItems:     cartItems,
		},
	}
}
