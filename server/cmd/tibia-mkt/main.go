package main

import (
	"log"
	"os"

	"github.com/adriein/tibia-mkt/internal/server"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if os.Getenv(constants.Env) != constants.Production {
		dotenvErr := godotenv.Load()

		if dotenvErr != nil {
			log.Fatal("Error loading .env file")
		}
	}

	checker := helper.NewEnvVarChecker(
		constants.DatabaseUser,
		constants.DatabasePassword,
		constants.DatabaseName,
		constants.ServerPort,
		constants.TibiaMktApiKey,
		constants.Env,
	)

	if envCheckerErr := checker.Check(); envCheckerErr != nil {
		log.Fatal(envCheckerErr.Error())
	}

	server.New(os.Getenv(constants.ServerPort))
}
