package api

import (
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var OrderRoutes = []shared.RouterSchema{
	{
		RouteMethod: shared.RouteMethodPost,
		Path:   		 "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler: 		 createOrder,
	},
	{
		RouteMethod: shared.RouteMethodPut,
		Path:   		 "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler: 		 updateOrder,
	},
	{
		RouteMethod: shared.RouteMethodGet,
		Path:   		 "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler: 		 getOrder,
	},
	{
		RouteMethod: shared.RouteMethodDelete,
		Path:   		 "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler: 		 cancelOrder,
	},
}