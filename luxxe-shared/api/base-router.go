package api

import (
	"github.com/gofiber/fiber/v2"

	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

func mergeMiddlewares(fMiddlewares []typings.FiberMiddleware, handler typings.FiberMiddleware) []func(*fiber.Ctx) error {
	// Convert the slice of FiberMiddleware to a slice of func(*fiber.Ctx) error
	var middlewares []func(*fiber.Ctx) error
	for _, mw := range fMiddlewares {
		middlewares = append(middlewares, mw)
	}

	// Append the handler function to the middlewares slice
	middlewares = append(middlewares, handler)
	return middlewares
}

func BaseRouter(router fiber.Router, schema []shared.RouterSchema) {
	for _, value := range schema {
		switch value.RouteMethod {
		case shared.RouteMethodGet:
			router.Get(value.Path, mergeMiddlewares(value.Middlewares, value.Handler)...)
		case shared.RouteMethodPost:
			router.Post(value.Path, mergeMiddlewares(value.Middlewares, value.Handler)...)
		case shared.RouteMethodPut:
			router.Put(value.Path, mergeMiddlewares(value.Middlewares, value.Handler)...)
		case shared.RouteMethodDelete:
			router.Delete(value.Path, mergeMiddlewares(value.Middlewares, value.Handler)...)
		}
	}
}
