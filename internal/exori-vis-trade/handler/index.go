package handler

import (
	"log/slog"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handle root")
}
