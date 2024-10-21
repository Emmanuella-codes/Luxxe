package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	repo_product "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type Products struct {
	ProductCount	int64								`json:"productCount"`
	Products			*[]entities.Product	`json:"products"`
}

func GetAllProductsPipe(ctx context.Context, dto *dtos.GetProductDTO) *shared.PipeRes[Products] {
	page := dto.Page

	products, productCount, _ := repo_product.ProductRepo.QueryAllProducts(ctx, page)

	return &shared.PipeRes[Products]{
		Success: true,
		Message: messages.SUCCESS_GET_PRODUCTS,
		Data:  	 &Products{
			ProductCount: productCount,
			Products:	 products,
		},
	}
}
