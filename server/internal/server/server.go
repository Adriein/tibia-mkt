package server

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/adriein/tibia-mkt/internal/detail"
	"github.com/adriein/tibia-mkt/internal/event"
	"github.com/adriein/tibia-mkt/internal/health"
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/internal/script"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/middleware"
	"github.com/adriein/tibia-mkt/pkg/statistics"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rotisserie/eris"
)

type TibiaMkt struct {
	database  *sql.DB
	router    *gin.RouterGroup
	validator *validator.Validate
}

func New(port string) *TibiaMkt {
	engine := gin.Default()
	router := engine.Group("/api/v1")

	router.Use(middleware.Error())

	app := &TibiaMkt{
		database:  initDatabase(),
		router:    router,
		validator: validator.New(),
	}

	app.routeSetup()

	if ginErr := engine.Run(port); ginErr != nil {
		err := eris.Wrap(ginErr, "Error starting HTTP server")

		log.Fatal(eris.ToString(err, true))
	}

	slog.Info("Starting the PingrateApiServer at " + port)

	return app
}

func initDatabase() *sql.DB {
	databaseDsn := fmt.Sprintf(
		"postgresql://%s:%s@localhost:5432/%s?sslmode=disable",
		os.Getenv(constants.DatabaseUser),
		os.Getenv(constants.DatabasePassword),
		os.Getenv(constants.DatabaseName),
	)

	database, dbConnErr := sql.Open("postgres", databaseDsn)

	if dbConnErr != nil {
		log.Fatal(dbConnErr.Error())
	}

	return database
}

func (t *TibiaMkt) routeSetup() {
	//HEALTH CHECK
	t.router.GET("/ping", health.NewController().Get())

	//SCRIPTS
	scriptController := t.getScriptController()

	t.router.GET("/scripts/prices", scriptController.SeedPrices())

	//PRICES
	priceController := t.getPriceController()

	t.router.GET("/prices", priceController.Get())

	//DETAIL
	detailController := t.getDetailController()

	t.router.GET("/details", detailController.Get())

	//EVENT
	eventController := t.getEventController()

	t.router.GET("/events", eventController.Get())
}

func (t *TibiaMkt) getPriceController() *price.Controller {
	repository := price.NewPgPriceRepository(t.database)
	service := price.NewService(repository)
	presenter := price.NewPresenter()

	return price.NewController(service, presenter)
}

func (t *TibiaMkt) getScriptController() *script.Controller {
	csvDataRepository := script.NewCsvSecuraPricesRepository()
	jsonDataRepository := script.NewJsonSecuraPricesRepository()
	pricesRepository := price.NewPgPriceRepository(t.database)

	service := script.NewService(csvDataRepository, jsonDataRepository, pricesRepository)

	return script.NewController(service)
}

func (t *TibiaMkt) getDetailController() *detail.Controller {
	priceRepository := price.NewPgPriceRepository(t.database)
	priceService := price.NewService(priceRepository)

	service := detail.NewService(statistics.New(), priceService)
	presenter := detail.NewPresenter()

	return detail.NewController(service, presenter)
}

func (t *TibiaMkt) getEventController() *event.Controller {
	repository := event.NewPgEventRepository(t.database)
	service := event.NewService(repository)

	presenter := event.NewPresenter()

	return event.NewController(service, presenter)
}
