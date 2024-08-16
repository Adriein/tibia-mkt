package main

import (
	"database/sql"
	"errors"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/handler"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/presenter"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/server"
	service2 "github.com/adriein/tibia-mkt/internal/tibia-mkt/service"
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/internal/trade-engine/trade-algorithm"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/middleware"
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

	cronMiddlewares := middleware.NewMiddlewareChain(
		middleware.NewAuthMiddleWare,
	)

	api.Route("GET /home", createHomeHandler(api, database))
	api.Route("GET /detail", createDetailHandler(api, database))
	api.Route("GET /kill-statistics-cron", cronMiddlewares.ApplyOn(createKillStatisticsHandler(api, database)))
	api.Route("GET /data-snapshot-cron", cronMiddlewares.ApplyOn(createDataSnapshotHandler(api, database)))

	api.Route("POST /trade-engine", tradeEngineHandler(api, database))
	api.Route("POST /seed", createSeedHandler(api, database))

	api.Start()

	defer database.Close()
}

func createHomeHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := helper.NewContainer(database, pgGoodRepository)

	homePresenter := presenter.NewHomePresenter(pgGoodRepository)

	factory, err := container.NewGoodRecordRepositoryFactory()

	if err != nil {
		log.Fatal(err.Error())
	}

	home := handler.NewHomeHandler(factory, homePresenter)

	return api.NewHandler(home.Handler)
}

func createSeedHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := helper.NewContainer(database, pgGoodRepository)

	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()

	factory, err := container.NewGoodRecordRepositoryFactory()

	if err != nil {
		log.Fatal(err.Error())
	}

	pgDataSnapshotRepository := repository.NewPgDataSnapshotRepository(database)
	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)

	prob := helper.NewProbHelper()

	detailService := service2.NewDetailService(
		pgGoodRepository,
		pgKillStatisticRepository,
		pgDataSnapshotRepository,
		factory,
		prob,
	)

	dataCron := service2.NewDataSnapshotCron(pgGoodRepository, pgDataSnapshotRepository, detailService)

	seederService := service2.NewSeeder(csvSecuraCogRepository, pgGoodRepository, container, dataCron)

	seed := handler.NewSeedHandler(seederService)

	return api.NewHandler(seed.Handler)
}

func tradeEngineHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := helper.NewContainer(database, pgGoodRepository)

	homePresenter := presenter.NewHomePresenter(pgGoodRepository)

	factory, err := container.NewGoodRecordRepositoryFactory()

	if err != nil {
		log.Fatal(err.Error())
	}

	config := trade_engine.NewConfig()
	prob := helper.NewProbHelper()

	algorithm := trade_algorithm.NewBestSellValueAlgorithm(config, prob)

	engine := trade_engine.NewTradeEngine[trade_algorithm.BestSellValueResult](factory, config, algorithm)

	engineHandler := handler.NewTradeEngineHandler(engine, homePresenter)

	return api.NewHandler(engineHandler.Handler)
}

func createDetailHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := helper.NewContainer(database, pgGoodRepository)

	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)
	pgDataSnapshotRepository := repository.NewPgDataSnapshotRepository(database)

	detailPresenter := presenter.NewDetailPresenter()

	factory, err := container.NewGoodRecordRepositoryFactory()

	if err != nil {
		log.Fatal(err.Error())
	}

	prob := helper.NewProbHelper()

	detailService := service2.NewDetailService(
		pgGoodRepository,
		pgKillStatisticRepository,
		pgDataSnapshotRepository,
		factory,
		prob,
	)

	detail := handler.NewDetailHandler(detailService, detailPresenter)

	return api.NewHandler(detail.Handler)
}

func createKillStatisticsHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)
	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)

	command := service2.NewKillStatisticsCron()

	killStatistics := handler.NewKillStatisticsHandler(command, pgGoodRepository, pgKillStatisticRepository)

	return api.NewHandler(killStatistics.Handler)
}

func createDataSnapshotHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)
	container := helper.NewContainer(database, pgGoodRepository)

	pgDataSnapshotRepository := repository.NewPgDataSnapshotRepository(database)
	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)

	factory, err := container.NewGoodRecordRepositoryFactory()

	if err != nil {
		log.Fatal(err.Error())
	}

	prob := helper.NewProbHelper()

	detailService := service2.NewDetailService(
		pgGoodRepository,
		pgKillStatisticRepository,
		pgDataSnapshotRepository,
		factory,
		prob,
	)

	command := service2.NewDataSnapshotCron(pgGoodRepository, pgDataSnapshotRepository, detailService)

	dataSnapshot := handler.NewDataSnapshotHandler(command)

	return api.NewHandler(dataSnapshot.Handler)
}
