package repository

import (
	"database/sql"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type PgTibiaCoinRepository struct {
	connection *sql.DB
	table string
}

func NewPgTibiaCoinRepository(connection *sql.DB) *PgTibiaCoinRepository {
	return &PgTibiaCoinRepository{
		connection: connection,
		table: "tibia_coin"
	}
}

func (r *PgTibiaCoinRepository) Find(criteria types.Criteria) ([]types.CogSku, error) {
	return nil, nil
}

func (r *PgTibiaCoinRepository) Save(entity types.CogSku) error {
	var command = `INSERT INTO tibia_coin (id, world, date, price, action_type) VALUES ($1, $2, $3, $4, $5)`

	r.connection.ExecContext()
	return nil
}
