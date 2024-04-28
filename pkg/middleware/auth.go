package middleware

import (
	"log/slog"
	"net/http"
)

func NewAuthMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Auth middleware")
		handler.ServeHTTP(w, r)
	})
}
