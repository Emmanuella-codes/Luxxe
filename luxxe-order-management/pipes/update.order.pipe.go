package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/dtos"
	cart_messages "github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	order_messages "github.com/Emmanuella-codes/Luxxe/luxxe-order-management/messages"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	order_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/order"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func UpdateOrderPipe(ctx context.Context, dto *dtos.UpdateOrderDTO) *shared.PipeRes[entities.OrderManagement] {
	userID, orderID, shippingAddress, phoneNumber := dto.UserID, dto.OrderID, dto.ShippingAddress, dto.PhoneNumber 

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

	_, err = order_repo.OrderRepo.QueryByID(ctx, orderID)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: order_messages.FAIL_GET_ORDER,
		}
	}

	updatedOrder := &entities.OrderManagement{
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
		Message: order_messages.SUCCESS_CREATE_ORDER,
		Data: order,
	}
}
