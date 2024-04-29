package handler

import (
	"encoding/json"
	"github.com/adriein/exori-vis-trade/pkg/types"
	"net/http"
)

type HomeResponse struct {
	SecuraTibiaCoin types.CogSku `json:"securaTibiaCoin"`
}

type HomeHandler struct {
	repository types.CogRepository
}

func NewHomeHandler(repository types.CogRepository) *HomeHandler {
	return &HomeHandler{repository: repository}
}

func (hh *HomeHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	data := &types.ServerResponse{
		Ok:   true,
		Data: nil,
	}

	bytes, err := json.Marshal(data)

	if err != nil {
		return &types.EvtError{
			Msg:      err.Error(),
			Function: "HomeHandler",
			File:     "home.go",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errs := w.Write(bytes)
	if errs != nil {
		return &types.EvtError{
			Msg:      errs.Error(),
			Function: "HomeHandler",
			File:     "home.go",
		}
	}

	return nil
}
