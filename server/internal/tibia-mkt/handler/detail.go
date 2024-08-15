package handler

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type DetailHandler struct {
	service   *service.DetailService
	presenter types.Presenter
}

func NewDetailHandler(
	service *service.DetailService,
	presenter types.Presenter,
) *DetailHandler {
	return &DetailHandler{
		service:   service,
		presenter: presenter,
	}
}

func (h *DetailHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	paramsMap := r.URL.Query()

	if !paramsMap.Has("item") {
		return types.ApiError{
			Msg:      constants.NoGoodSearchParamProvided,
			Function: "HomeHandler",
			File:     "home.go",
			Domain:   true,
		}
	}

	cog := paramsMap["item"][0]

	result, serviceErr := h.service.Execute(cog)

	if serviceErr != nil {
		return serviceErr
	}

	response, presenterErr := h.presenter.Format(result)

	if presenterErr != nil {
		return presenterErr
	}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
