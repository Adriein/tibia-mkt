package repository

import (
	"database/sql"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type PgTibiaCoinRepository struct {
	connection *sql.DB
}

func NewPgTibiaCoinRepository(connection *sql.DB) *PgTibiaCoinRepository {
	return &PgTibiaCoinRepository{
		connection: connection,
	}
}

func (r *PgTibiaCoinRepository) Find(criteria types.Criteria) ([]types.CogSku, error) {
	return nil, nil
}

func (r *PgTibiaCoinRepository) Save(entity types.CogSku) error {
	var query = `INSERT INTO tibia_coin (id, world, date, price, action_type) VALUES ($1, $2, $3, $4, $5)`

	result, err := r.connection.Exec(query, entity.Id, entity.World, entity.Date, entity.Price, entity.Action)

	fmt.Println(result)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     "pg-tibia-coin-repository.go",
		}
	}

	return nil
}
