package api

import (
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var AuthRoutes = []shared.RouterSchema{
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/signup",
		Handler:     signUpUser,
	},
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/signin",
		Handler:     signInUser,
	},
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/send-otp",
		Handler:     sendOtp,
	},
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/reset-password",
		Handler:     resetUserPasswordByEmail,
	},
	{
		RouteMethod: shared.RouteMethodPost,
		Path:        "/reset-password-id",
		Middlewares: []typings.FiberMiddleware{services.BaseAuthToken, services.IsAnyUserMiddleware},
		Handler:     resetUserPasswordByID,
	},
	{
		RouteMethod: shared.RouteMethodGet,
		Path:        "/verify-email",
		Handler:     verifyEmail,
	},
}
