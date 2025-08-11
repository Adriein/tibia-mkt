package event

import (
	"database/sql"
	"strings"
	"time"

	"github.com/rotisserie/eris"
)

type EventRepository interface {
	FindByWorldAndGood(world string, good string) ([]*Event, error)
	Save(Event *Event) error
}

type PgEventRepository struct {
	connection *sql.DB
}

func NewPgEventRepository(connection *sql.DB) *PgEventRepository {
	return &PgEventRepository{
		connection: connection,
	}
}

func (r *PgEventRepository) FindByWorldAndGood(worldName string, good string) ([]*Event, error) {
	statement, err := r.connection.Prepare("SELECT * FROM events WHERE world = $1 AND good_name = $2 ORDER BY occurred_at DESC;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		id          string
		name        string
		good_name   string
		world       string
		description string
		occurred_at string
	)

	rows, err := statement.Query(worldName, good)

	defer rows.Close()

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var results []*Event

	for rows.Next() {
		scanErr := rows.Scan(&id, &name, &good_name, &world, &description, &occurred_at)

		if scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		occurredAt, occurredAtTimeParseErr := time.Parse(time.RFC3339Nano, occurred_at)

		if occurredAtTimeParseErr != nil {
			return nil, eris.New(occurredAtTimeParseErr.Error())
		}

		results = append(results, &Event{
			Id:          id,
			Name:        name,
			GoodName:    good_name,
			World:       world,
			Description: description,
			OccurredAt:  occurredAt,
		})
	}

	return results, nil
}

func (r *PgEventRepository) Save(Event *Event) error {
	var b strings.Builder
	b.WriteString("INSERT INTO Events (")
	b.WriteString("id, name, good_name, world, description, occurred_at")
	b.WriteString(") VALUES ($1, $2, $3, $4, $5, $6)")

	var query = b.String()

	_, err := r.connection.Exec(
		query,
		Event.Id,
		Event.Name,
		Event.GoodName,
		Event.GoodName,
		Event.World,
		Event.Description,
		Event.OccurredAt.Format(time.DateTime),
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}
