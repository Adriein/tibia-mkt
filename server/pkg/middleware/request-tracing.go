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

		initialTrace := fmt.Sprintf(
			"Received request: Method=%s URI=%s UserAgent=%s",
			r.Method,
			r.RequestURI,
			r.UserAgent(),
		)

		slog.Info(initialTrace)

		handler.ServeHTTP(w, r)

		elapsed := time.Since(start)

		trace := fmt.Sprintf(
			"Method=%s URI=%s UserAgent=%s Time=%dms",
			r.Method,
			r.RequestURI,
			r.UserAgent(),
			elapsed.Milliseconds(),
		)

		slog.Info(trace)
	})
}
