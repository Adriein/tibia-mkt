package handler

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type HomeHandler struct {
	repoFactory *helper.RepositoryFactory
	presenter   types.Presenter
}

func NewHomeHandler(factory *helper.RepositoryFactory, presenter types.Presenter) *HomeHandler {
	return &HomeHandler{
		repoFactory: factory,
		presenter:   presenter,
	}
}

func (h *HomeHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	var repositoryResponseMatrix [][]types.GoodRecord

	paramsMap := r.URL.Query()

	if !paramsMap.Has("item") {
		return types.ApiError{
			Msg:      constants.NoGoodSearchParamProvided,
			Function: "HomeHandler",
			File:     "home.go",
			Domain:   true,
		}
	}

	params := paramsMap["item"]

	for _, good := range params {
		repository := h.repoFactory.Get(good)

		var filters []types.Filter

		filters = append(filters, types.Filter{Name: "world", Operand: constants.Equal, Value: "Secura"})

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

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
