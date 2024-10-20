package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-product/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	product_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type EmptyStruct struct{}

func DeleteProductPipe(ctx context.Context, dto *dtos.DeleteProductDTO) *shared.PipeRes[EmptyStruct] {
	productID := dto.ProductID
	_, err := product_repo.ProductRepo.QueryByID(ctx, productID)

	product_repo.ProductRepo.DeleteProduct(ctx, productID)

	if err != nil {
		return &shared.PipeRes[EmptyStruct]{
			Success: false,
			Message: messages.FAIL_UPDATE_PRODUCT,
			Data:    &EmptyStruct{},
		}
	}

	return &shared.PipeRes[EmptyStruct]{
		Success: true,
		Message: messages.SUCCESS_UPDATE_PRODUCT,
		Data:    &EmptyStruct{},
	}
}
