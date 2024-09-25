package api

import shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"

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
}
