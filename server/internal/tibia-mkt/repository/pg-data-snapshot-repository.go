package repository

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type PgDataSnapshotRepository struct {
	connection  *sql.DB
	transformer *service.CriteriaToSqlService
}

func NewPgDataSnapshotRepository(connection *sql.DB) *PgDataSnapshotRepository {
	transformer := service.NewCriteriaToSqlService("data_snapshot_cron")

	return &PgDataSnapshotRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgDataSnapshotRepository) Find(criteria types.Criteria) ([]types.DataSnapshot, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, types.ApiError{
			Msg:      err.Error(),
			Function: "Find -> r.transformer.Transform()",
			File:     "pg-data-snapshot-repository.go",
		}
	}

	rows, queryErr := r.connection.Query(query)

	if queryErr != nil {
		return nil, types.ApiError{
			Msg:      queryErr.Error(),
			Function: "Find -> r.connection.Query()",
			File:     "pg-data-snapshot-repository.go",
		}
	}

	defer rows.Close()

	var (
		id            string
		cog           string
		std_deviation float64
		mean          int
		total_droped  int
		executed_by   string
		created_at    string
		updated_at    string
	)

	var results []types.DataSnapshot

	for rows.Next() {
		if scanErr := rows.Scan(
			&id,
			&cog,
			&std_deviation,
			&mean,
			&total_droped,
			&executed_by,
			&created_at,
			&updated_at,
		); scanErr != nil {
			return nil, types.ApiError{
				Msg:      scanErr.Error(),
				Function: "Find -> rows.Scan()",
				File:     "pg-data-snapshot-repository.go",
			}
		}

		results = append(results, types.DataSnapshot{
			Id:           id,
			Cog:          cog,
			StdDeviation: std_deviation,
			Mean:         mean,
			TotalDropped: total_droped,
			ExecutedBy:   executed_by,
			CreatedAt:    created_at,
			UpdatedAt:    updated_at,
		})
	}

	return results, nil
}

func (r *PgDataSnapshotRepository) FindOne(criteria types.Criteria) (types.DataSnapshot, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return types.DataSnapshot{}, types.ApiError{
			Msg:      err.Error(),
			Function: "FindOne -> r.transformer.Transform()",
			File:     "pg-data-snapshot-repository.go",
		}
	}

	var (
		id            string
		cog           string
		std_deviation float64
		mean          int
		total_droped  int
		executed_by   string
		created_at    string
		updated_at    string
	)

	if scanErr := r.connection.QueryRow(query).Scan(
		&id,
		&cog,
		&std_deviation,
		&mean,
		&total_droped,
		&executed_by,
		&created_at,
		&updated_at,
	); scanErr != nil {
		return types.DataSnapshot{}, types.ApiError{
			Msg:      scanErr.Error(),
			Function: "FindOne -> rows.Scan()",
			File:     "pg-data-snapshot-repository.go",
			Values:   []string{query},
		}
	}

	return types.DataSnapshot{
		Id:           id,
		Cog:          cog,
		StdDeviation: std_deviation,
		Mean:         mean,
		TotalDropped: total_droped,
		ExecutedBy:   executed_by,
		CreatedAt:    created_at,
		UpdatedAt:    updated_at,
	}, nil
}

func (r *PgDataSnapshotRepository) Save(entity types.DataSnapshot) error {
	var query = `INSERT INTO data_snapshot_cron (id, cog, std_deviation, mean, total_droped, executed_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.connection.Exec(
		query,
		entity.Id,
		entity.Cog,
		entity.StdDeviation,
		entity.Mean,
		entity.TotalDropped,
		entity.ExecutedBy,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     "pg-data-snapshot-repository.go",
		}
	}

	return nil
}
