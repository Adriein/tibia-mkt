package main

import (
	"database/sql"
	"errors"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/handler"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/presenter"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/server"
	"github.com/adriein/tibia-mkt/pkg/middleware"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

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

	api.Route("/home", createHomeHandler(api, database))
	api.Route("/foo", fooMiddlewares.ApplyOn(api.NewHandler(handler.FooHandler)))
	api.Route("/seed", createSeedHandler(api, database))

	api.Start()

	defer database.Close()
}

func createHomeHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgSecuraTibiaCoinCogRepository := repository.NewPgTibiaCoinRepository(database)
	pgSecuraHoneycombCogRepository := repository.NewPgHoneycombRepository(database)

	pgCogRepository := repository.NewPgCogRepository(database)

	homePresenter := presenter.NewHomePresenter(pgCogRepository)

	var repositories []types.CogRepository
	repositories = append(repositories, pgSecuraTibiaCoinCogRepository, pgSecuraHoneycombCogRepository)

	factory := service.NewRepositoryFactory(repositories)

	home := handler.NewHomeHandler(factory, homePresenter)

	return api.NewHandler(home.Handler)
}

func createSeedHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()
	pgCogRepository := repository.NewPgHoneycombRepository(database)

	seed := handler.NewSeedHandler(csvSecuraCogRepository, pgCogRepository)

	return api.NewHandler(seed.Handler)
}
