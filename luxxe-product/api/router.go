package api

import (
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var ProductRoutes = []shared.RouterSchema{
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     createProduct,
	},
	{
		RouteMethod: shared.RouteMethodDelete,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     deleteProduct,
	},
	{
		RouteMethod: shared.RouteMethodPut,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     updateProduct,
	},
	{
		RouteMethod: shared.RouteMethodGet,
		Path:        "/all",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     getProducts,
	},
	{
		RouteMethod: shared.RouteMethodGet,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     getProductsByCategory,
	},
}
