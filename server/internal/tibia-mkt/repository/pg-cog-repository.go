package repository

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type PgCogRepository struct {
	connection  *sql.DB
	transformer *service.CriteriaToSqlService
}

func NewPgCogRepository(connection *sql.DB) *PgCogRepository {
	transformer := service.NewCriteriaToSqlService("cog")

	return &PgCogRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgCogRepository) FindOne(criteria types.Criteria) (types.Cog, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return types.Cog{}, types.ApiError{
			Msg:      err.Error(),
			Function: "FindOne -> r.transformer.Transform()",
			File:     "pg-cog-repository.go",
		}
	}

	rows, queryErr := r.connection.Query(query)

	defer rows.Close()

	if queryErr != nil {
		return types.Cog{}, types.ApiError{
			Msg:      queryErr.Error(),
			Function: "FindOne -> r.connection.Query()",
			File:     "pg-cog-repository.go",
		}
	}

	var (
		id         string
		name       string
		link       string
		created_at string
		updated_at string
	)

	var results []types.Cog

	for rows.Next() {
		if scanErr := rows.Scan(&id, &name, &link, &created_at, &updated_at); scanErr != nil {
			return types.Cog{}, types.ApiError{
				Msg:      scanErr.Error(),
				Function: "FindOne -> rows.Scan()",
				File:     "pg-cog-repository.go",
			}
		}

		parsedCreatedAt, createdAtParseErr := time.Parse(time.DateTime, created_at)

		if createdAtParseErr != nil {
			return types.Cog{}, types.ApiError{
				Msg:      createdAtParseErr.Error(),
				Function: "FindOne -> time.Parse()",
				File:     "pg-cog-repository.go",
			}
		}

		parsedUpdatedAt, udpatedAtParseErr := time.Parse(time.DateTime, updated_at)

		if udpatedAtParseErr != nil {
			return types.Cog{}, types.ApiError{
				Msg:      udpatedAtParseErr.Error(),
				Function: "FindOne -> time.Parse()",
				File:     "pg-cog-repository.go",
			}
		}

		results = append(results, types.Cog{
			Id:        id,
			Name:      name,
			Link:      link,
			CreatedAt: parsedCreatedAt,
			UpdatedAt: parsedUpdatedAt,
		})
	}

	return results[0], nil
}
