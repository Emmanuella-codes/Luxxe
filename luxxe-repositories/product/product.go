package product

import (
	"context"

	"github.com/go-kit/log"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) (*entities.Product, error)
	QueryAllProducts(ctx context.Context, page int) (*[]entities.Product, int64, error)
	UpdateProductByID(ctx context.Context, 
		productID string, 
		name string, 
		description string,
	  price float64,
		category entities.ProductCategories,
		productImage string,
		productInfo string,
		quantity int) (*entities.Product, error)
	QueryByID(ctx context.Context, productID string) (*entities.Product, error)
	QueryProductsByCategory(ctx context.Context, category entities.ProductCategories, page int) (*[]entities.Product, int64, error)
	DeleteProduct(ctx context.Context, productID string)
}

var ProductRepo ProductRepository

func InitProductRepo(logger *log.Logger) {
	ProductRepo = newMgRepository(logger)
}
