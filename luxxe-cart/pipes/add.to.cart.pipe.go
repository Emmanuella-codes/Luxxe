package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	user_messages"github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func AddToCartPipe(ctx context.Context, dto *dtos.AddToCartDTO) *shared.PipeRes[entities.Cart] {
	userIDStr, productID, quantity := dto.UserID, dto.ProductID, dto.Quantity

	_, error := repo_user.UserRepo.QueryByID(ctx, userIDStr)
	if error != nil {
		return &shared.PipeRes[entities.Cart]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	cart, error := cart_repo.CartRepo.AddToCart(ctx, userIDStr, productID, quantity)
	if error != nil {
		return &shared.PipeRes[entities.Cart]{
			Success: false,
			Message: messages.FAIL_ADD_TO_CART,
		}
	}

	return &shared.PipeRes[entities.Cart]{
		Success: true,
		Message: messages.SUCCESS_ADD_TO_CART,
		Data:    cart,
	}
}
