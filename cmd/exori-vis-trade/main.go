package main

import (
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/handler"
	"github.com/adriein/exori-vis-trade/internal/exori-vis-trade/server"
	"os"
)

func main() {
	api, err := server.New(":8080")

	api.Route("/index", handler.HomeHandler)

	if err != nil {
		os.Exit(1)
	}

	api.Start()
}
