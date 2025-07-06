package handler

import (
	"github.com/adriein/tibia-mkt/internal/presenter"
	"github.com/adriein/tibia-mkt/internal/service"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type SearchGoodHandler struct {
	service   *service.SearchGoodService
	presenter *presenter.SearchGoodPresenter
}

func NewSearchGoodHandler(
	service *service.SearchGoodService,
	presenter *presenter.SearchGoodPresenter,
) *SearchGoodHandler {
	return &SearchGoodHandler{
		service:   service,
		presenter: presenter,
	}
}

func (h *SearchGoodHandler) Handler(w http.ResponseWriter, _ *http.Request) error {

	goods, goodServiceErr := h.service.Execute()

	if goodServiceErr != nil {
		return goodServiceErr
	}

	response, presenterErr := h.presenter.Format(goods)

	if presenterErr != nil {
		return presenterErr
	}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
