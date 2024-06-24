package repository

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type PgTibiaCoinRepository struct {
	connection  *sql.DB
	transformer *service.CriteriaToSqlService
	name        string
}

func NewPgTibiaCoinRepository(connection *sql.DB) *PgTibiaCoinRepository {
	transformer := service.NewCriteriaToSqlService("tibia_coin")

	return &PgTibiaCoinRepository{
		connection:  connection,
		transformer: transformer,
		name:        constants.TibiaCoinEntity,
	}
}

func (r *PgTibiaCoinRepository) Find(criteria types.Criteria) ([]types.CogSku, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, types.ApiError{
			Msg:      err.Error(),
			Function: "Find -> r.transformer.Transform()",
			File:     "pg-tibia-coin-repository.go",
		}
	}

	rows, queryErr := r.connection.Query(query)

	defer rows.Close()

	if queryErr != nil {
		return nil, types.ApiError{
			Msg:      queryErr.Error(),
			Function: "Find -> r.connection.Query()",
			File:     "pg-tibia-coin-repository.go",
		}
	}

	var (
		id         string
		world      string
		date       string
		buy_price  float64
		sell_price float64
	)

	var results []types.CogSku

	for rows.Next() {
		if scanErr := rows.Scan(&id, &world, &date, &buy_price, &sell_price); scanErr != nil {
			return nil, types.ApiError{
				Msg:      scanErr.Error(),
				Function: "Find -> rows.Scan()",
				File:     "pg-tibia-coin-repository.go",
			}
		}

		buyPrice := int(buy_price)
		sellPrice := int(sell_price)

		parsedDate, timeParseErr := time.Parse(time.DateOnly, date)

		if timeParseErr != nil {
			return nil, types.ApiError{
				Msg:      timeParseErr.Error(),
				Function: "Find -> time.Parse()",
				File:     "pg-tibia-coin-repository.go",
			}
		}

		results = append(results, types.CogSku{
			Id:        id,
			ItemName:  constants.TibiaCoinEntity,
			Date:      parsedDate,
			BuyPrice:  buyPrice,
			SellPrice: sellPrice,
			World:     world,
		})
	}

	return results, nil
}

func (r *PgTibiaCoinRepository) Save(entity types.CogSku) error {
	var query = `INSERT INTO tibia_coin (id, world, date, buy_price, sell_price) VALUES ($1, $2, $3, $4, $5)`

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
			File:     "pg-tibia-coin-repository.go",
		}
	}

	return nil
}

func (r *PgTibiaCoinRepository) EntityName() string {
	return r.name
}
