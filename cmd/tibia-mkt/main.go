package main

import (
	"database/sql"
	"errors"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/handler"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/presenter"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/server"
	"github.com/adriein/tibia-mkt/pkg/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	api, newServerErr := server.New(":4000")

	if newServerErr != nil {
		log.Fatal(newServerErr.Error())
	}

	databaseDSN, existEnv := os.LookupEnv("DATABASE_URL")

	if !existEnv {
		noEnvVarError := errors.New("DATABASE_URL is not set")

		log.Fatal(noEnvVarError.Error())
	}

	database, dbConnErr := sql.Open("postgres", databaseDSN)

	if dbConnErr != nil {
		log.Fatal(dbConnErr.Error())
	}

	fooMiddlewares := middleware.NewMiddlewareChain(
		middleware.NewAuthMiddleWare,
	)

	api.Route("/home", createHomeHandler(api))
	api.Route("/foo", fooMiddlewares.ApplyOn(api.NewHandler(handler.FooHandler)))

	api.Start()

	defer database.Close()
}

func createHomeHandler(api *server.TibiaMktApiServer) http.HandlerFunc {
	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()
	homePresenter := presenter.NewHomePresenter()

	home := handler.NewHomeHandler(csvSecuraCogRepository, homePresenter)

	return api.NewHandler(home.Handler)
}
