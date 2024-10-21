package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	product_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func CreateProductPipe(ctx context.Context, dto *dtos.CreateProductDTO) *shared.PipeRes[entities.Product] {
	name, desciption, category, price, quantity, productImage, productInfo := dto.Name, 
		dto.Description, dto.Category, dto.Price, dto.Quantity, dto.ProductImage, dto.ProductInfo

	product := &entities.Product{
		Name: 				name,
		Description: 	desciption,
		Category: 		category,
		Price: 				price,
		Quantity: 		quantity,
		ProductImage: productImage,
		ProductInfo: 	productInfo,
	}
	
	newProduct, err := product_repo.ProductRepo.Create(ctx, product)
	if err != nil {
		return &shared.PipeRes[entities.Product]{
			Success: false,
			Message: messages.FAIL_CREATE_PRODUCT,
		}
	}

	return &shared.PipeRes[entities.Product]{
		Success: true,
		Message: messages.SUCCESS_CREATE_PRODUCT,
		Data: 	 newProduct,
	}
}
