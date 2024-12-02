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

func UpdateCartItemPipe(ctx context.Context, dto *dtos.UpdateCartItemDTO) *shared.PipeRes[entities.Cart] {
	userIDStr, productID, quantity, price := dto.UserID, dto.ProductID, dto.Quantity, dto.Price

	_, error := repo_user.UserRepo.QueryByID(ctx, userIDStr)
	if error != nil {
		return &shared.PipeRes[entities.Cart]{
			Success: false,
			Message: user_messages.NOT_FOUND_USER,
		}
	}

	cart, error := cart_repo.CartRepo.UpdateCartItem(ctx, userIDStr, productID, quantity, price)
	if error != nil {
		return &shared.PipeRes[entities.Cart]{
			Success: false,
			Message: messages.FAIL_UPDATE_CART,
		}
	}

	return &shared.PipeRes[entities.Cart]{
		Success: true,
		Message: messages.SUCCESS_UPDATE_CART,
		Data: 	 cart,
	}
}
