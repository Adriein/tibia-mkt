package server

import (
	"database/sql"
	"fmt"
	"github.com/adriein/tibia-mkt/internal/health"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rotisserie/eris"
	"log"
	"log/slog"
	"os"
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
		pingrateErr := eris.Wrap(ginErr, "Error starting HTTP server")

		log.Fatal(eris.ToString(pingrateErr, true))
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

func (p *TibiaMkt) routeSetup() {
	//HEALTH CHECK
	p.router.GET("/ping", health.NewController().Get())
}
