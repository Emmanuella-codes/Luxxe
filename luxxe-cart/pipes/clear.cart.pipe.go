package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type EmptyStruct struct{}

func ClearCartPipe(ctx context.Context, dto *dtos.ClearCartDTO) *shared.PipeRes[EmptyStruct] {
	userIDStr := dto.UserID

	_, err := repo_user.UserRepo.QueryByID(ctx, userIDStr)
	if err != nil {
		return &shared.PipeRes[EmptyStruct]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	err = cart_repo.CartRepo.ClearCart(ctx, userIDStr)
	if err != nil {
		return &shared.PipeRes[EmptyStruct]{
			Success: false,
			Message: messages.FAIL_CLEAR_CART,
			Data:    &EmptyStruct{},
		}
	}

	return &shared.PipeRes[EmptyStruct]{
		Success: true,
		Message: messages.SUCCESS_CLEAR_CART,
		Data:    &EmptyStruct{},
	}
}