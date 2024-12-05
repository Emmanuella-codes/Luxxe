package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/dtos"
	order_messages "github.com/Emmanuella-codes/Luxxe/luxxe-order-management/messages"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	order_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/order"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func GetOrderPipe(ctx context.Context, dto *dtos.GetOrderDTO) *shared.PipeRes[entities.OrderManagement] {
	userID := dto.UserID

	_, err := repo_user.UserRepo.QueryByID(ctx, userID)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	order, err := order_repo.OrderRepo.GetOrder(ctx, userID)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: order_messages.FAIL_GET_ORDER,
		}
	}

	return &shared.PipeRes[entities.OrderManagement]{
		Success: true,
		Message: order_messages.SUCCESS_GET_ORDER,
		Data: 	 order,
	}
}
