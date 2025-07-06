package price

import "database/sql"

type PriceRepository interface {
	FindByNameAndWorld(email string) (*Price, error)
}

type PgPriceRepository struct {
	connection *sql.DB
}

func NewPgPriceRepository(connection *sql.DB) *PgPriceRepository {
	return &PgPriceRepository{
		connection: connection,
	}
}
