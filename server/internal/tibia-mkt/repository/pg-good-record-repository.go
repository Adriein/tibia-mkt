package repository

import (
	"database/sql"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type PgGoodRecordRepository struct {
	connection  *sql.DB
	transformer *helper.CriteriaToSqlService
	name        string
}

func NewPgGoodRecordRepository(connection *sql.DB, name string) *PgGoodRecordRepository {
	transformer := helper.NewCriteriaToSqlService(name)

	return &PgGoodRecordRepository{
		connection:  connection,
		transformer: transformer,
		name:        name,
	}
}

func (r *PgGoodRecordRepository) Find(criteria types.Criteria) ([]types.GoodRecord, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, types.ApiError{
			Msg:      err.Error(),
			Function: "Find -> r.transformer.Transform()",
			File:     fmt.Sprintf("pg-%s-repository.go", r.name),
		}
	}

	rows, queryErr := r.connection.Query(query)

	if queryErr != nil {
		return nil, types.ApiError{
			Msg:      queryErr.Error(),
			Function: "Find -> r.connection.Query()",
			File:     fmt.Sprintf("pg-%s-repository.go", r.name),
		}
	}

	defer rows.Close()

	var (
		id         string
		world      string
		date       string
		buy_price  float64
		sell_price float64
	)

	var results []types.GoodRecord

	for rows.Next() {
		if scanErr := rows.Scan(&id, &world, &date, &buy_price, &sell_price); scanErr != nil {
			return nil, types.ApiError{
				Msg:      scanErr.Error(),
				Function: "Find -> rows.Scan()",
				File:     fmt.Sprintf("pg-%s-repository.go", r.name),
			}
		}

		buyPrice := int(buy_price)
		sellPrice := int(sell_price)

		parsedDate, timeParseErr := time.Parse(time.DateOnly, date)

		if timeParseErr != nil {
			return nil, types.ApiError{
				Msg:      timeParseErr.Error(),
				Function: "Find -> time.Parse()",
				File:     fmt.Sprintf("pg-%s-repository.go", r.name),
			}
		}

		results = append(results, types.GoodRecord{
			Id:        id,
			ItemName:  helper.SnakeToCamel(r.name),
			Date:      parsedDate,
			BuyPrice:  buyPrice,
			SellPrice: sellPrice,
			World:     world,
		})
	}

	return results, nil
}

func (r *PgGoodRecordRepository) Save(entity types.GoodRecord) error {
	var query = fmt.Sprintf("INSERT INTO %s (id, world, date, buy_price, sell_price) VALUES ($1, $2, $3, $4, $5)", r.name)

	_, err := r.connection.Exec(
		query,
		entity.Id,
		entity.World,
		entity.Date.Format(time.DateOnly),
		entity.BuyPrice,
		entity.SellPrice,
	)

	if err != nil {
		return types.ApiError{
			Msg:      err.Error(),
			Function: "Save -> r.connection.Exec()",
			File:     fmt.Sprintf("pg-%s-repository.go", r.name),
		}
	}

	return nil
}

func (r *PgGoodRecordRepository) GoodName() string {
	return r.name
}
