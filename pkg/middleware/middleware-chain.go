package middleware

import (
	"github.com/adriein/exori-vis-trade/pkg/types"
	"net/http"
)

type Chain struct {
	middlewares []types.Middleware
}

func New(initialMiddlewares ...types.Middleware) *Chain {
	return &Chain{
		middlewares: initialMiddlewares,
	}
}

func (c *Chain) Apply(next http.Handler) http.Handler {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		next = c.middlewares[i](next)
	}

	return next
}
