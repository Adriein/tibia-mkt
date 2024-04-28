package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

func NewRequestTracingMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("%s: %s", r.Method, r.RequestURI))
		handler.ServeHTTP(w, r)
	})
}
