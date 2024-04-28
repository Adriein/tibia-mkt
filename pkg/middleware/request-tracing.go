package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func NewRequestTracingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		elapsed := time.Since(start)

		trace := fmt.Sprintf(
			"Method=%s URI=%s UserAgent=%s Time=%dns",
			r.Method,
			r.RequestURI,
			r.UserAgent(),
			elapsed,
		)

		slog.Info(trace)
	})
}
