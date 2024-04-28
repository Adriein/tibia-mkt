package main

import (
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/handler"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/server"
	"os"
)

func main() {
	api, err := server.New(":4000")

	api.Route("/home", api.NewHandler(handler.HomeHandler))
	api.Route("/foo", api.NewHandler(handler.FooHandler))

	if err != nil {
		os.Exit(1)
	}

	api.Start()
}
