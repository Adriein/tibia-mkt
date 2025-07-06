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

func (r *PgGoodRepository) FindByName(name string) (*Good, error) {
	statement, err := r.connection.Prepare("SELECT * FROM ti_good WHERE name = $1;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		ti_id         string
		ti_name       string
		ti_wiki_link  string
		ti_creatures  []byte
		ti_created_at string
		ti_updated_at string
	)

	if scanErr := statement.QueryRow(name).Scan(
		&ti_id,
		&ti_name,
		&ti_wiki_link,
		&ti_creatures,
		&ti_created_at,
		&ti_updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(GoodNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	parsedCreatedAt, createdAtParseErr := time.Parse(time.DateTime, ti_created_at)

	if createdAtParseErr != nil {
		return nil, eris.New(createdAtParseErr.Error())
	}

	parsedUpdatedAt, updatedAtParseErr := time.Parse(time.DateTime, ti_updated_at)

	if updatedAtParseErr != nil {
		return nil, eris.New(updatedAtParseErr.Error())
	}

	decodedCreatures, decodeError := helper.JsonDecode[[]Creature](ti_creatures)

	if decodeError != nil {
		return nil, eris.New(decodeError.Error())
	}

	return &Good{
		Id:        ti_id,
		Name:      ti_name,
		WikiLink:  ti_wiki_link,
		Creatures: decodedCreatures,
		CreatedAt: parsedCreatedAt,
		UpdatedAt: parsedUpdatedAt,
	}, nil
}
