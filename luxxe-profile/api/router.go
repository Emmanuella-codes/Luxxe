package api

import (
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var UserRoutes = []shared.RouterSchema{
	{
		RouteMethod: shared.RouteMethodGet,
		Path:        "/",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     getUserByAccountToken,
	},
}

