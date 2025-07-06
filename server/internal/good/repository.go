package good

import "database/sql"

type GoodRepository interface {
	FindByNameAndWorld(email string) (*Good, error)
}

type PgGoodRepository struct {
	connection *sql.DB
}

func NewPgGoogleTokenRepository(connection *sql.DB) *PgGoodRepository {
	return &PgGoodRepository{
		connection: connection,
	}
}
