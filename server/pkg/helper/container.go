package helper

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type Container struct {
	database       *sql.DB
	goodRepository types.Repository[types.Good]
}

func NewContainer(database *sql.DB, goodRepository types.Repository[types.Good]) *Container {
	return &Container{
		database:       database,
		goodRepository: goodRepository,
	}
}

func (c *Container) NewGoodRecordRepositoryFactory() (*RepositoryFactory, error) {
	goods, err := c.goodRepository.Find(types.Criteria{Filters: make([]types.Filter, 0)})

	if err != nil {
		return nil, err
	}

	var repositories = make([]types.GoodRecordRepository, len(goods))

	for index, good := range goods {
		repositories[index] = repository.NewPgGoodRecordRepository(c.database, CamelToSnake(good.Name))
	}

	return NewRepositoryFactory(repositories), nil
}
