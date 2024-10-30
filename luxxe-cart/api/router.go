package api

import (
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var CartRoutes = []shared.RouterSchema{
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     addToCart,
	},
	{
		RouteMethod: shared.RouteMethodPut,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     updateCart,
	},
	{
		RouteMethod: shared.RouteMethodGet,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     getCart,
	},
	{
		RouteMethod: shared.RouteMethodDelete,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     clearCart,
	},
	{
		RouteMethod: shared.RouteMethodDelete,
		Path:        "/item",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     removeItemFromCart,
	},
}
