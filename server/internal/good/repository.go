package good

import (
	"database/sql"
	"errors"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/rotisserie/eris"
	"time"
)

type GoodRepository interface {
	FindByName(name string) (*Good, error)
}

type PgGoodRepository struct {
	connection *sql.DB
}

func NewPgGoogleTokenRepository(connection *sql.DB) *PgGoodRepository {
	return &PgGoodRepository{
		connection: connection,
	}
}

func (r *PgGoodRepository) FindByName(goodName string) (*Good, error) {
	statement, err := r.connection.Prepare("SELECT * FROM good WHERE name = $1;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		id         string
		name       string
		wiki_link  string
		creatures  []byte
		created_at string
		updated_at string
	)

	if scanErr := statement.QueryRow(goodName).Scan(
		&id,
		&name,
		&wiki_link,
		&creatures,
		&created_at,
		&updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(GoodNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	parsedCreatedAt, createdAtParseErr := time.Parse(time.DateTime, created_at)

	if createdAtParseErr != nil {
		return nil, eris.New(createdAtParseErr.Error())
	}

	parsedUpdatedAt, updatedAtParseErr := time.Parse(time.DateTime, updated_at)

	if updatedAtParseErr != nil {
		return nil, eris.New(updatedAtParseErr.Error())
	}

	decodedCreatures, decodeError := helper.JsonDecode[[]Creature](creatures)

	if decodeError != nil {
		return nil, eris.New(decodeError.Error())
	}

	return &Good{
		Id:        id,
		Name:      name,
		WikiLink:  wiki_link,
		Creatures: decodedCreatures,
		CreatedAt: parsedCreatedAt,
		UpdatedAt: parsedUpdatedAt,
	}, nil
}
