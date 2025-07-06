package repository

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type PgGoodRepository struct {
	connection  *sql.DB
	transformer *helper.CriteriaToSqlService
}

func NewPgGoodRepository(connection *sql.DB) *PgGoodRepository {
	transformer := helper.NewCriteriaToSqlService("good")

	return &PgGoodRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgGoodRepository) Find(criteria types.Criteria) ([]types.Good, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, types.ApiError{
			Msg:      err.Error(),
			Function: "Find -> r.transformer.Transform()",
			File:     "pg-good-repository.go",
		}
	}

	rows, queryErr := r.connection.Query(query)

	if queryErr != nil {
		return nil, types.ApiError{
			Msg:      queryErr.Error(),
			Function: "Find -> r.connection.Query()",
			File:     "pg-good-repository.go",
		}
	}

	defer rows.Close()

	var (
		id         string
		name       string
		link       string
		creatures  []byte
		created_at string
		updated_at string
	)

	var results []types.Good

	for rows.Next() {
		if scanErr := rows.Scan(&id, &name, &link, &creatures, &created_at, &updated_at); scanErr != nil {
			return nil, types.ApiError{
				Msg:      scanErr.Error(),
				Function: "Find -> rows.Scan()",
				File:     "pg-good-repository.go",
			}
		}

		parsedCreatedAt, createdAtParseErr := time.Parse(time.DateTime, created_at)

		if createdAtParseErr != nil {
			return nil, types.ApiError{
				Msg:      createdAtParseErr.Error(),
				Function: "Find -> time.Parse()",
				File:     "pg-good-repository.go",
			}
		}

		parsedUpdatedAt, updatedAtParseErr := time.Parse(time.DateTime, updated_at)

		if updatedAtParseErr != nil {
			return nil, types.ApiError{
				Msg:      updatedAtParseErr.Error(),
				Function: "Find -> time.Parse()",
				File:     "pg-good-repository.go",
			}
		}

		decodedCreatures, decodeError := helper.JsonDecode[[]types.GoodDrop](creatures)

		if decodeError != nil {
			return nil, types.ApiError{
				Msg:      decodeError.Error(),
				Function: "Find -> helper.JsonDecode()",
				File:     "pg-good-repository.go",
			}
		}

		results = append(results, types.Good{
			Id:        id,
			Name:      name,
			Link:      link,
			Drop:      decodedCreatures,
			CreatedAt: parsedCreatedAt,
			UpdatedAt: parsedUpdatedAt,
		})
	}

	return results, nil
}

func (r *PgGoodRepository) FindOne(criteria types.Criteria) (types.Good, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return types.Good{}, types.ApiError{
			Msg:      err.Error(),
			Function: "FindOne -> r.transformer.Transform()",
			File:     "pg-good-repository.go",
		}
	}

	var (
		id         string
		name       string
		link       string
		creatures  []byte
		created_at string
		updated_at string
	)

	if scanErr := r.connection.QueryRow(query).Scan(&id, &name, &link, &creatures, &created_at, &updated_at); scanErr != nil {
		return types.Good{}, types.ApiError{
			Msg:      scanErr.Error(),
			Function: "FindOne -> rows.Scan()",
			File:     "pg-good-repository.go",
			Values:   []string{query},
		}
	}

	parsedCreatedAt, createdAtParseErr := time.Parse(time.DateTime, created_at)

	if createdAtParseErr != nil {
		return types.Good{}, types.ApiError{
			Msg:      createdAtParseErr.Error(),
			Function: "FindOne -> time.Parse()",
			File:     "pg-good-repository.go",
		}
	}

	parsedUpdatedAt, updatedAtParseErr := time.Parse(time.DateTime, updated_at)

	if updatedAtParseErr != nil {
		return types.Good{}, types.ApiError{
			Msg:      updatedAtParseErr.Error(),
			Function: "FindOne -> time.Parse()",
			File:     "pg-good-repository.go",
		}
	}

	decodedCreatures, decodeError := helper.JsonDecode[[]types.GoodDrop](creatures)

	if decodeError != nil {
		return types.Good{}, types.ApiError{
			Msg:      decodeError.Error(),
			Function: "FindOne -> helper.JsonDecode()",
			File:     "pg-good-repository.go",
		}
	}

	return types.Good{
		Id:        id,
		Name:      name,
		Link:      link,
		Drop:      decodedCreatures,
		CreatedAt: parsedCreatedAt,
		UpdatedAt: parsedUpdatedAt,
	}, nil
}

func (r *PgGoodRepository) Save(entity types.Good) error {
	var query = `INSERT INTO good (id, name, link, creatures, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	encodedCreatures, jsonEncodeErr := helper.JsonEncode(entity.Drop)

	if jsonEncodeErr != nil {
		return types.ApiError{
			Msg:      jsonEncodeErr.Error(),
			Function: "Save -> helper.JsonEncode()",
			File:     "pg-good-repository.go",
		}
	}

	_, err := r.connection.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Link,
		encodedCreatures,
		entity.CreatedAt.Format(time.DateTime),
		entity.UpdatedAt.Format(time.DateTime),
	)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     "pg-good-repository.go",
		}
	}

	return nil
}
