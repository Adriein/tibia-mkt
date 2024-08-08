package repository

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type PgKillStatisticRepository struct {
	connection  *sql.DB
	transformer *service.CriteriaToSqlService
}

func NewPgKillStatisticRepository(connection *sql.DB) *PgKillStatisticRepository {
	transformer := service.NewCriteriaToSqlService("kill_statistic_cron")

	return &PgKillStatisticRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgKillStatisticRepository) Find(criteria types.Criteria) ([]types.KillStatistic, error) {
	return nil, types.ApiError{
		Msg:      "Not implemented yet",
		Function: "Find -> r.transformer.Transform()",
		File:     "pg-kill-statistic-repository.go",
	}
}

func (r *PgKillStatisticRepository) FindOne(criteria types.Criteria) (types.KillStatistic, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return types.KillStatistic{}, types.ApiError{
			Msg:      err.Error(),
			Function: "FindOne -> r.transformer.Transform()",
			File:     "pg-kill-statistic-repository.go",
		}
	}

	var (
		id            string
		creature_name string
		amount_killed int
		drop_rate     float64
		executed_by   string
		created_at    string
		updated_at    string
	)

	if scanErr := r.connection.QueryRow(query).Scan(
		&id,
		&creature_name,
		&amount_killed,
		&drop_rate,
		&executed_by,
		&created_at,
		&updated_at,
	); scanErr != nil {
		return types.KillStatistic{}, types.ApiError{
			Msg:      scanErr.Error(),
			Function: "FindOne -> rows.Scan()",
			File:     "pg-kill-statistic-repository.go",
			Values:   []string{query},
		}
	}

	return types.KillStatistic{
		Id:           id,
		CreatureName: creature_name,
		AmountKilled: amount_killed,
		DropRate:     drop_rate,
		ExecutedBy:   executed_by,
		CreatedAt:    created_at,
		UpdatedAt:    updated_at,
	}, nil
}

func (r *PgKillStatisticRepository) Save(entity types.KillStatistic) error {
	var query = `INSERT INTO kill_statistic_cron (id, creature_name, amount_killed, drop_rate, executed_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.connection.Exec(
		query,
		entity.Id,
		entity.CreatureName,
		entity.AmountKilled,
		entity.DropRate,
		entity.ExecutedBy,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     "pg-kill-statistic-repository.go",
		}
	}

	return nil
}
