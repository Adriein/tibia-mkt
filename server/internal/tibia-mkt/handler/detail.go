package handler

import (
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type DetailHandler struct {
	repoFactory *service.RepositoryFactory
	presenter   types.Presenter
}

func NewDetailHandler(factory *service.RepositoryFactory, presenter types.Presenter) *DetailHandler {
	return &DetailHandler{
		repoFactory: factory,
		presenter:   presenter,
	}
}

func (h *DetailHandler) Handler(w http.ResponseWriter, _ *http.Request) error {
	return nil
}
