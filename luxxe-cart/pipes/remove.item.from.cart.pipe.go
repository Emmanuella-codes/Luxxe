package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func RemoveItemFromCartPipe(ctx context.Context, dto *dtos.RemoveItemFromCartDTO) *shared.PipeRes[entities.Cart] {
	userIDStr, productID := dto.UserID, dto.ProductID

	_, error := repo_user.UserRepo.QueryByID(ctx, userIDStr)
	if error != nil {
		return &shared.PipeRes[entities.Cart]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	cart, error := cart_repo.CartRepo.RemoveFromCart(ctx, userIDStr, productID)
	if error != nil {
		return &shared.PipeRes[entities.Cart]{
			Success: false,
			Message: messages.FAIL_REMOVE_ITEM_FROM_CART,
		}
	}

	return &shared.PipeRes[entities.Cart]{
		Success: true,
		Message: messages.SUCCESS_REMOVE_ITEM_FROM_CART,
		Data: 	 cart,
	}
}
