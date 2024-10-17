package main

import (
	"database/sql"
	"fmt"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/handler"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/presenter"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/server"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/service"
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/internal/trade-engine/trade-algorithm"
	"github.com/adriein/tibia-mkt/pkg"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv(constants.Env) != constants.Production {
		dotenvErr := godotenv.Load()

		if dotenvErr != nil {
			log.Fatal("Error loading .env file")
		}
	}

	checker := helper.NewEnvVarChecker(
		constants.DatabaseIp,
		constants.DatabaseUser,
		constants.DatabasePassword,
		constants.DatabaseName,
		constants.ServerPort,
		constants.TibiaMktApiKey,
	)

	if envCheckerErr := checker.Check(); envCheckerErr != nil {
		log.Fatal(envCheckerErr.Error())
	}

	api, newServerErr := server.New(os.Getenv(constants.ServerPort))

	if newServerErr != nil {
		log.Fatal(newServerErr.Error())
	}

	databaseDsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv(constants.DatabaseUser),
		os.Getenv(constants.DatabasePassword),
		os.Getenv(constants.DatabaseIp),
		os.Getenv(constants.DatabaseName),
	)

	database, dbConnErr := sql.Open("postgres", databaseDsn)

	if dbConnErr != nil {
		log.Fatal(dbConnErr.Error())
	}

	cronMiddlewares := middleware.NewMiddlewareChain(
		middleware.NewAuthMiddleWare,
	)

	api.Route("GET /home", createHomeHandler(api, database))
	api.Route("GET /detail", createDetailHandler(api, database))
	api.Route("GET /goods", createSearchGoodsHandler(api, database))

	api.Route("GET /kill-statistics-cron", cronMiddlewares.ApplyOn(createKillStatisticsHandler(api, database)))
	api.Route("GET /data-snapshot-cron", cronMiddlewares.ApplyOn(createDataSnapshotHandler(api, database)))
	api.Route("POST /data-ingestion", cronMiddlewares.ApplyOn(createDataIngestionHandler(api, database)))

	api.Route("POST /trade-engine", tradeEngineHandler(api, database))
	api.Route("POST /seed", createSeedHandler(api, database))

	api.Start()

	defer database.Close()
}

func createHomeHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := pkg.NewContainer(database)

	homePresenter := presenter.NewHomePresenter(pgGoodRepository)

	factory := container.NewGoodRecordRepositoryFactory()

	home := handler.NewHomeHandler(factory, homePresenter)

	return api.NewHandler(home.Handler)
}

func createSeedHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := pkg.NewContainer(database)

	csvSecuraCogRepository := repository.NewCsvSecuraCogRepository()

	factory := container.NewGoodRecordRepositoryFactory()

	pgDataSnapshotRepository := repository.NewPgDataSnapshotRepository(database)
	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)

	prob := helper.NewProbHelper()

	detailService := service.NewDetailService(
		pgGoodRepository,
		pgKillStatisticRepository,
		pgDataSnapshotRepository,
		factory,
		prob,
	)

	dataCron := service.NewDataSnapshotService(pgGoodRepository, pgDataSnapshotRepository, detailService)

	seederService := service.NewSeederService(csvSecuraCogRepository, pgGoodRepository, dataCron, factory)

	seed := handler.NewSeedHandler(seederService)

	return api.NewHandler(seed.Handler)
}

func tradeEngineHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := pkg.NewContainer(database)

	homePresenter := presenter.NewHomePresenter(pgGoodRepository)

	factory := container.NewGoodRecordRepositoryFactory()

	config := trade_engine.NewConfig()
	prob := helper.NewProbHelper()

	algorithm := trade_algorithm.NewBestSellValueAlgorithm(config, prob)

	engine := trade_engine.NewTradeEngine[trade_algorithm.BestSellValueResult](factory, config, algorithm)

	engineHandler := handler.NewTradeEngineHandler(engine, homePresenter)

	return api.NewHandler(engineHandler.Handler)
}

func createDetailHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	container := pkg.NewContainer(database)

	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)
	pgDataSnapshotRepository := repository.NewPgDataSnapshotRepository(database)

	detailPresenter := presenter.NewDetailPresenter()

	factory := container.NewGoodRecordRepositoryFactory()

	prob := helper.NewProbHelper()

	detailService := service.NewDetailService(
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

	command := service.NewKillStatisticsCron()

	killStatistics := handler.NewKillStatisticsHandler(command, pgGoodRepository, pgKillStatisticRepository)

	return api.NewHandler(killStatistics.Handler)
}

func createDataSnapshotHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)
	container := pkg.NewContainer(database)

	pgDataSnapshotRepository := repository.NewPgDataSnapshotRepository(database)
	pgKillStatisticRepository := repository.NewPgKillStatisticRepository(database)

	factory := container.NewGoodRecordRepositoryFactory()

	prob := helper.NewProbHelper()

	detailService := service.NewDetailService(
		pgGoodRepository,
		pgKillStatisticRepository,
		pgDataSnapshotRepository,
		factory,
		prob,
	)

	command := service.NewDataSnapshotService(pgGoodRepository, pgDataSnapshotRepository, detailService)

	dataSnapshot := handler.NewDataSnapshotHandler(command)

	return api.NewHandler(dataSnapshot.Handler)
}

func createSearchGoodsHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	pgGoodRepository := repository.NewPgGoodRepository(database)

	goodsPresenter := presenter.NewSearchGoodPresenter()

	goodService := service.NewSearchGoodService(
		pgGoodRepository,
	)

	searchGood := handler.NewSearchGoodHandler(goodService, goodsPresenter)

	return api.NewHandler(searchGood.Handler)
}

func createDataIngestionHandler(api *server.TibiaMktApiServer, database *sql.DB) http.HandlerFunc {
	container := pkg.NewContainer(database)

	factory := container.NewGoodRecordRepositoryFactory()

	dataIngestionService := service.NewDataIngestionService(factory)

	dataIngestion := handler.NewDataIngestionHandler(dataIngestionService)

	return api.NewHandler(dataIngestion.Handler)
}
