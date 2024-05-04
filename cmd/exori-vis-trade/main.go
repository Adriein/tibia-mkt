package main

import (
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/handler"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/presenter"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/repository"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/server"
	"github.com/adriein/exori-vis-trade/pkg/middleware"
	"net/http"
	"os"
)

func main() {
	api, err := server.New(":4000")

	fooMiddlewares := middleware.NewMiddlewareChain(
		middleware.NewAuthMiddleWare,
	)

	api.Route("/home", createHomeHandler(api))
	api.Route("/foo", fooMiddlewares.ApplyOn(api.NewHandler(handler.FooHandler)))

	if err != nil {
		os.Exit(1)
	}

	api.Start()
}

func createHomeHandler(api *server.ExoriVisTradeApiServer) http.HandlerFunc {
	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()
	homePresenter := presenter.NewHomePresenter()

	home := handler.NewHomeHandler(csvSecuraCogRepository, homePresenter)

	return api.NewHandler(home.Handler)
}
