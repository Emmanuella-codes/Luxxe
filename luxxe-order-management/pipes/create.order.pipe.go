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

func CreateOrderPipe(ctx context.Context, dto *dtos.CreateOrderDTO) *shared.PipeRes[entities.OrderManagement] {
	userIDStr,  shippingAddress, phoneNumber := dto.UserID, dto.ShippingAddress, dto.PhoneNumber

	_, error := repo_user.UserRepo.QueryByID(ctx, userIDStr)
	if error != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	cart, err := cart_repo.CartRepo.QueryByUserID(ctx, userIDStr)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: cart_messages.NOT_FOUND_CART,
		}
	}

	orderObj := &entities.OrderManagement{
		UserID:       		cart.UserID,
		CartID:       		cart.ID,
		ShippingAddress: 	shippingAddress,
		PhoneNumber: 			phoneNumber,
	}

	order, err := order_repo.OrderRepo.Create(ctx, orderObj)
	if err != nil {
		return &shared.PipeRes[entities.OrderManagement]{
			Success: false,
			Message: order_messages.FAIL_CREATE_ORDER,
		}
	}

	return &shared.PipeRes[entities.OrderManagement]{
		Success: true,
		Message: order_messages.SUCCESS_CREATE_ORDER,
		Data: order,
	}
}
