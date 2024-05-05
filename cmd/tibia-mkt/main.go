package main

import (
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/handler"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/presenter"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/server"
	"github.com/adriein/tibia-mkt/pkg/middleware"
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

func createHomeHandler(api *server.TibiaMktApiServer) http.HandlerFunc {
	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()
	homePresenter := presenter.NewHomePresenter()

	home := handler.NewHomeHandler(csvSecuraCogRepository, homePresenter)

	return api.NewHandler(home.Handler)
}
