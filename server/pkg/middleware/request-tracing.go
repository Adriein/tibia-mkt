package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

func NewRequestTracingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		id := uuid.New()

		r.Header.Add("traceId", id.String())

		initialTrace := fmt.Sprintf(
			"Received request: Method=%s URI=%s UserAgent=%s TraceId=%s",
			r.Method,
			r.RequestURI,
			r.UserAgent(),
			id.String(),
		)

		slog.Info(initialTrace)

		handler.ServeHTTP(w, r)

		elapsed := time.Since(start)

		trace := fmt.Sprintf(
			"Method=%s URI=%s UserAgent=%s Time=%dms TraceId=%s",
			r.Method,
			r.RequestURI,
			r.UserAgent(),
			elapsed.Milliseconds(),
			id.String(),
		)

		slog.Info(trace)
	})
}
