package api

import (
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var TransactionsRoutes = []shared.RouterSchema{
	{
		RouteMethod: shared.RouteMethodPost,
		Path:   		 "/verify",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler: 		 verifyTransaction,
	},
	{
		RouteMethod: shared.RouteMethodPost,
		Path:   		 "/webhook/paystack",
		Handler: 		 paystackWebHook,
	},
	{
		RouteMethod: shared.RouteMethodGet,
		Path:   		 "/all",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler: 		 getAllInitiatorsTransactions,
	},
}
