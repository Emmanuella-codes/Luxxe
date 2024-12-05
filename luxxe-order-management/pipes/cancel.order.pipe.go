package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/dtos"
	order_messages "github.com/Emmanuella-codes/Luxxe/luxxe-order-management/messages"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	order_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/order"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type EmptyStruct struct{}

func CancelOrderPipe(ctx context.Context, dto *dtos.CancelOrderDTO) *shared.PipeRes[EmptyStruct] {
	userIDStr := dto.UserID

	_, err := repo_user.UserRepo.QueryByID(ctx, userIDStr)
	if err != nil {
		return &shared.PipeRes[EmptyStruct]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	_, err = order_repo.OrderRepo.QueryByUserID(ctx, userIDStr)
	if err != nil {
		return &shared.PipeRes[EmptyStruct]{
			Success: false,
			Message: order_messages.FAIL_GET_ORDER,
		}
	}

	err = order_repo.OrderRepo.CancelOrder(ctx, userIDStr)
	if err != nil {
		return &shared.PipeRes[EmptyStruct]{
			Success: false,
			Message: order_messages.FAIL_CANCEL_ORDER,
			Data:    &EmptyStruct{},
		}
	}

	return &shared.PipeRes[EmptyStruct]{
		Success: true,
		Message: order_messages.SUCCESS_CANCEL_ORDER,
		Data: 	 &EmptyStruct{},
	}
}