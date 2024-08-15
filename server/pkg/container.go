package pkg

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"unicode"
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

func (c *Container) NewGoodRecordRepositoryFactory() (*service.RepositoryFactory, error) {
	goods, err := c.goodRepository.Find(types.Criteria{Filters: make([]types.Filter, 0)})

	if err != nil {
		return nil, err
	}

	var repositories = make([]types.GoodRecordRepository, len(goods))

	for index, good := range goods {
		repositories[index] = repository.NewPgGoodRecordRepository(c.database, c.camelToSnake(good.Name))
	}

	return service.NewRepositoryFactory(repositories), nil
}

func (c *Container) camelToSnake(str string) string {
	var result []rune

	for i, char := range str {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
	}

	return string(result)
}
