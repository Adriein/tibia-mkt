package handler

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
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
			Msg:      constants.NoCogSearchParamProvided,
			Function: "HomeHandler",
			File:     "home.go",
			Domain:   true,
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

	response, presenterErr := h.presenter.Format(repositoryResponseMatrix)

	if presenterErr != nil {
		return presenterErr
	}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
