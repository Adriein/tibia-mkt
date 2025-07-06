package pkg

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/internal/repository"
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
	goods := [4]string{
		constants.TibiaCoinEntity,
		constants.HoneycombEntity,
		constants.SwamplingWoodEntity,
		constants.BrokenShamanicStaffEntity,
	}
	var repositories = make([]types.GoodRecordRepository, len(goods))

	for index, good := range goods {
		repositories[index] = repository.NewPgGoodRecordRepository(c.database, helper.CamelToSnake(good))
	}

	return helper.NewRepositoryFactory(repositories)
}
