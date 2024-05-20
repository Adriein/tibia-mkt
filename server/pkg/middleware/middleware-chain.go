package middleware

import (
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type MiddlewareChain struct {
	middlewares []types.Middleware
}

func NewMiddlewareChain(initialMiddlewares ...types.Middleware) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: initialMiddlewares,
	}
}

func (c *MiddlewareChain) ApplyOn(next http.Handler) http.Handler {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		next = c.middlewares[i](next)
	}

	return next
}
