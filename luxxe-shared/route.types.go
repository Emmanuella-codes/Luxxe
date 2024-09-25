package shared

import "github.com/Emmanuella-codes/Luxxe/typings"

type RouteMethod string

const RouteMethodGet RouteMethod = "get"
const RouteMethodPost RouteMethod = "post"
const RouteMethodPut RouteMethod = "put"
const RouteMethodDelete RouteMethod = "delete"

type RouterSchema struct {
	RouteMethod RouteMethod
	Path        string
	Middlewares []typings.FiberMiddleware
	Handler     typings.FiberMiddleware
}
