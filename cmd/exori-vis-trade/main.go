package main

import (
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/handler"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/repository"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/server"
	"github.com/adriein/exori-vis-trade/pkg/middleware"
	"os"
)

func main() {
	api, err := server.New(":4000")

	fooMiddlewares := middleware.NewMiddlewareChain(
		middleware.NewAuthMiddleWare,
	)

	repo := repository.NewCsvSecuraCogRepository()

	home := handler.NewHomeHandler(repo)

	api.Route("/home", api.NewHandler(home.Handler))
	api.Route("/foo", fooMiddlewares.ApplyOn(api.NewHandler(handler.FooHandler)))

	if err != nil {
		os.Exit(1)
	}

	api.Start()
}
