package handler

import (
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type HomeResponse struct {
	SecuraTibiaCoin types.CogSku `json:"securaTibiaCoin"`
}

type HomeHandler struct {
	repoFactory *service.RepositoryFactory
	presenter   types.Presenter
}

func NewHomeHandler(factory *service.RepositoryFactory, presenter types.Presenter) *HomeHandler {
	return &HomeHandler{
		repoFactory: factory,
		presenter:   presenter,
	}
}

func (h *HomeHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	var repositoryResponseMatrix [][]types.CogSku

	paramsMap := r.URL.Query()

	if !paramsMap.Has("item") {
		return types.ApiError{
			Msg:      "No cog search params provided",
			Function: "HomeHandler",
			File:     "home.go",
		}
	}

	params := paramsMap["item"]

	for _, cog := range params {
		repository := h.repoFactory.Get(cog)

		var filters []types.Filter

		filters = append(filters, types.Filter{Name: "world", Operand: "=", Value: "Secura"})

		results, repositoryErr := repository.Find(types.Criteria{Filters: filters})

		if repositoryErr != nil {
			return repositoryErr
		}

		if len(results) == 0 {
			continue
		}

		repositoryResponseMatrix = append(repositoryResponseMatrix, results)
	}

	bytes, presenterErr := h.presenter.Format(repositoryResponseMatrix)

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
