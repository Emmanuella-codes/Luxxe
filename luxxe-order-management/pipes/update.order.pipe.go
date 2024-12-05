package pipes

import (
	"context"

	cart_messages "github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/dtos"
	order_messages "github.com/Emmanuella-codes/Luxxe/luxxe-order-management/messages"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	order_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/order"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func UpdateOrderPipe(ctx context.Context, dto *dtos.UpdateOrderDTO) *shared.PipeRes[entities.OrderManagement] {
	userID, shippingAddress, phoneNumber := dto.UserID, dto.ShippingAddress, dto.PhoneNumber 

	_, err := repo_user.UserRepo.QueryByID(ctx, userID)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	cart, err := cart_repo.CartRepo.QueryByUserID(ctx, userID)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: cart_messages.NOT_FOUND_CART,
		}
	}

	orderMgt, err := order_repo.OrderRepo.QueryByUserID(ctx, userID)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: order_messages.FAIL_GET_ORDER,
		}
	}

	updatedOrder := &entities.OrderManagement{
		ID:          		 orderMgt.ID,
		UserID:          cart.UserID,
		ShippingAddress: shippingAddress,
		PhoneNumber:   	 phoneNumber,
		CartTotal:       cart.TotalAmount,
	}

	order, err := order_repo.OrderRepo.UpdateOrder(ctx, updatedOrder)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: order_messages.FAIL_UPDATE_ORDER,
		}
	}

	return &shared.PipeRes[entities.OrderManagement]{
		Success: true,
		Message: order_messages.SUCCESS_UPDATE_ORDER,
		Data: order,
	}
}
