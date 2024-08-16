package pkg

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
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

func (c *Container) NewGoodRecordRepositoryFactory() *helper.RepositoryFactory {
	goods := [2]string{constants.TibiaCoinEntity, constants.HoneycombEntity}
	var repositories = make([]types.GoodRecordRepository, len(goods))

	for index, good := range goods {
		repositories[index] = repository.NewPgGoodRecordRepository(c.database, helper.CamelToSnake(good))
	}

	return helper.NewRepositoryFactory(repositories)
}
