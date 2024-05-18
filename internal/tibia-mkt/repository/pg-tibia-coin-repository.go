package repository

import (
	"database/sql"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type PgTibiaCoinRepository struct {
	connection  *sql.DB
	transformer *service.CriteriaToSqlService
}

func NewPgTibiaCoinRepository(connection *sql.DB) *PgTibiaCoinRepository {
	transformer := service.NewCriteriaToSqlService("tibia_coin")

	return &PgTibiaCoinRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgTibiaCoinRepository) Find(criteria types.Criteria) ([]types.CogSku, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, types.ApiError{
			Msg:      err.Error(),
			Function: "Find -> r.transformer.Transform()",
			File:     "pg-tibia-coin-repository.go",
		}
	}

	rows, queryErr := r.connection.Query(query)

	fmt.Println(rows)

	if queryErr != nil {
		return nil, types.ApiError{
			Msg:      queryErr.Error(),
			Function: "Find -> r.connection.Query()",
			File:     "pg-tibia-coin-repository.go",
		}
	}

	return nil, nil
}

func (r *PgTibiaCoinRepository) Save(entity types.CogSku) error {
	var query = `INSERT INTO tibia_coin (id, world, date, price, action_type) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.connection.Exec(query, entity.Id, entity.World, entity.Date, entity.Price, entity.Action)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     "pg-tibia-coin-repository.go",
		}
	}

	return nil
}
