package handler

import (
	"github.com/adriein/tibia-mkt/pkg/types"
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

func (h *HomeHandler) Handler(w http.ResponseWriter, _ *http.Request) error {
	var filters []types.Filter

	filters = append(filters, types.Filter{Name: "world", Operand: "=", Value: "Secura"})

	results, repositoryErr := h.repository.Find(types.Criteria{Filters: filters})

	if repositoryErr != nil {
		return repositoryErr
	}

	bytes, presenterErr := h.presenter.Format(results)

	if presenterErr != nil {
		return presenterErr
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
