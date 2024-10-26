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

func ClearCartPipe(ctx context.Context, dto *dtos.ClearCartDTO) *shared.PipeRes[entities.Cart] {
	userIDStr := dto.UserID

	_, error := repo_user.UserRepo.QueryByID(ctx, userIDStr)

	cart_repo.CartRepo.ClearCart(ctx, userIDStr)

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