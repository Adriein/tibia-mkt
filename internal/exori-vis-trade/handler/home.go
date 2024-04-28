package handler

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)

	return nil
}
