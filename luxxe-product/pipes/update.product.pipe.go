package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	product_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func UpdateProductPipe(ctx context.Context, dto *dtos.UpdateProductDTO) *shared.PipeRes[entities.Product] {
	productIDStr, name, description, category, price, productImage, productInfo, quantity := dto.ProductID,
		dto.Name, dto.Description, dto.Category, dto.Price, dto.ProductImage, dto.ProductInfo, dto.Quantity

	product, err := product_repo.ProductRepo.UpdateProductByID(ctx, productIDStr, name, description, price, category, productImage, productInfo, quantity)

	if err != nil {
		return &shared.PipeRes[entities.Product]{
			Success: false,
			Message: messages.FAIL_UPDATE_PRODUCT,
		}
	}

	return &shared.PipeRes[entities.Product]{
		Success: true,
		Message: messages.SUCCESS_UPDATE_PRODUCT,
		Data:    product,
	}
}
