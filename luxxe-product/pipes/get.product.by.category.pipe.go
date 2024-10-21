package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	repo_product "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type ProductsByCategory struct {
	ProductCount	int64								`json:"productCount"`
	Products			*[]entities.Product	`json:"products"`
}

func GetProductsByCategoryPipe(ctx context.Context, dto *dtos.GetProductByCategoryDTO) *shared.PipeRes[ProductsByCategory] {
	category, page := dto.Category, dto.Page

	products, productCount, _ := repo_product.ProductRepo.QueryProductsByCategory(ctx, category, page)

	return &shared.PipeRes[ProductsByCategory]{
		Success: true,
		Message: messages.SUCCESS_GET_PRODUCTS,
		Data:  	 &ProductsByCategory{
			ProductCount: productCount,
			Products:	 products,
		},
	}
}
