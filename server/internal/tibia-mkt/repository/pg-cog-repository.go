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

	var (
		id         string
		name       string
		link       string
		created_at string
		updated_at string
	)

	if scanErr := r.connection.QueryRow(query).Scan(&id, &name, &link, &created_at, &updated_at); scanErr != nil {
		return types.Cog{}, types.ApiError{
			Msg:      scanErr.Error(),
			Function: "FindOne -> rows.Scan()",
			File:     "pg-cog-repository.go",
			Values:   []string{query},
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

	parsedUpdatedAt, updatedAtParseErr := time.Parse(time.DateTime, updated_at)

	if updatedAtParseErr != nil {
		return types.Cog{}, types.ApiError{
			Msg:      updatedAtParseErr.Error(),
			Function: "FindOne -> time.Parse()",
			File:     "pg-cog-repository.go",
		}
	}

	return types.Cog{
		Id:        id,
		Name:      name,
		Link:      link,
		CreatedAt: parsedCreatedAt,
		UpdatedAt: parsedUpdatedAt,
	}, nil
}

func (r *PgCogRepository) Save(entity types.Cog) error {
	var query = `INSERT INTO cog (id, name, link, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.connection.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Link,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     "pg-cog-repository.go",
		}
	}

	return nil
}
