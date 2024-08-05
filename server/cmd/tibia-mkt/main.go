package main

import (
	"database/sql"
	"errors"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/handler"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/presenter"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/server"
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/internal/trade-engine/trade-algorithm"
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

	api.Route("GET /home", createHomeHandler(api, database))
	api.Route("GET /detail", createDetailHandler(api, database))
	api.Route("POST /trade-engine", tradeEngineHandler(api, database))
	api.Route("GET /foo", fooMiddlewares.ApplyOn(api.NewHandler(handler.FooHandler)))
	api.Route("POST /seed", createSeedHandler(api, database))

	api.Start()

	defer database.Close()
}

func createHomeHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgSecuraTibiaCoinCogRepository := repository.NewPgTibiaCoinRepository(database)
	pgSecuraHoneycombCogRepository := repository.NewPgHoneycombRepository(database)

	pgCogRepository := repository.NewPgCogRepository(database)

	homePresenter := presenter.NewHomePresenter(pgCogRepository)

	repositories := []types.CogRepository{pgSecuraTibiaCoinCogRepository, pgSecuraHoneycombCogRepository}

	factory := service.NewRepositoryFactory(repositories)

	home := handler.NewHomeHandler(factory, homePresenter)

	return api.NewHandler(home.Handler)
}

func createSeedHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()
	pgSecuraTibiaCoinCogRepository := repository.NewPgTibiaCoinRepository(database)
	pgSecuraHoneycombCogRepository := repository.NewPgHoneycombRepository(database)
	pgCogRepository := repository.NewPgCogRepository(database)

	repositories := []types.CogRepository{pgSecuraTibiaCoinCogRepository, pgSecuraHoneycombCogRepository}

	factory := service.NewRepositoryFactory(repositories)

	seed := handler.NewSeedHandler(csvSecuraCogRepository, factory, pgCogRepository)

	return api.NewHandler(seed.Handler)
}

func tradeEngineHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgSecuraTibiaCoinCogRepository := repository.NewPgTibiaCoinRepository(database)
	pgSecuraHoneycombCogRepository := repository.NewPgHoneycombRepository(database)

	pgCogRepository := repository.NewPgCogRepository(database)

	homePresenter := presenter.NewHomePresenter(pgCogRepository)

	repositories := []types.CogRepository{pgSecuraTibiaCoinCogRepository, pgSecuraHoneycombCogRepository}

	factory := service.NewRepositoryFactory(repositories)

	config := trade_engine.NewConfig()
	prob := service.NewProbHelper()

	algorithm := trade_algorithm.NewBestSellValueAlgorithm(config, prob)

	engine := trade_engine.NewTradeEngine[trade_algorithm.BestSellValueResult](factory, config, algorithm)

	engineHandler := handler.NewTradeEngineHandler(engine, homePresenter)

	return api.NewHandler(engineHandler.Handler)
}

func createDetailHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgSecuraTibiaCoinCogRepository := repository.NewPgTibiaCoinRepository(database)
	pgSecuraHoneycombCogRepository := repository.NewPgHoneycombRepository(database)

	pgCogRepository := repository.NewPgCogRepository(database)

	homePresenter := presenter.NewDetailPresenter()

	repositories := []types.CogRepository{pgSecuraTibiaCoinCogRepository, pgSecuraHoneycombCogRepository}

	factory := service.NewRepositoryFactory(repositories)

	detail := handler.NewDetailHandler(pgCogRepository, factory, homePresenter)

	return api.NewHandler(detail.Handler)
}
