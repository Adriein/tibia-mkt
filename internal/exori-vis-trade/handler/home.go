package handler

import (
	"log/slog"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handle root")
}
