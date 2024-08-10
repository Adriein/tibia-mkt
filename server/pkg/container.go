package pkg

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type Container struct {
	database *sql.DB
}

func NewContainer(database *sql.DB) *Container {
	return &Container{
		database: database,
	}
}

func (c *Container) NewCogSkuRepositoryFactory() *service.RepositoryFactory {
	pgSecuraTibiaCoinCogRepository := repository.NewPgTibiaCoinRepository(c.database)
	pgSecuraHoneycombCogRepository := repository.NewPgHoneycombRepository(c.database)

	repositories := []types.CogRepository{pgSecuraTibiaCoinCogRepository, pgSecuraHoneycombCogRepository}

	return service.NewRepositoryFactory(repositories)
}
