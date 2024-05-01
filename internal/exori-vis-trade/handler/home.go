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
	presenter  types.Presenter
}

func NewHomeHandler(repository types.CogRepository, presenter types.Presenter) *HomeHandler {
	return &HomeHandler{
		repository: repository,
		presenter:  presenter,
	}
}

func (h *HomeHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	results, repositoryErr := h.repository.Find(types.Criteria{})

	if repositoryErr != nil {
		return repositoryErr
	}

	data := &types.ServerResponse{
		Ok:   true,
		Data: results,
	}

	bytes, jsonErr := json.Marshal(data)

	if jsonErr != nil {
		return types.ApiError{
			Msg:      jsonErr.Error(),
			Function: "HomeHandler",
			File:     "home.go",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, writeErr := w.Write(bytes)

	if writeErr != nil {
		return types.ApiError{
			Msg:      writeErr.Error(),
			Function: "HomeHandler",
			File:     "home.go",
		}
	}

	return nil
}
