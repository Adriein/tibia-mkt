package price

import (
	"database/sql"
	"github.com/rotisserie/eris"
	"time"
)

type PriceRepository interface {
	FindByNameAndWorld(world string, good string) ([]*Price, error)
}

type PgPriceRepository struct {
	connection *sql.DB
}

func NewPgPriceRepository(connection *sql.DB) *PgPriceRepository {
	return &PgPriceRepository{
		connection: connection,
	}
}

func (r *PgPriceRepository) FindByNameAndWorld(worldName string, good string) ([]*Price, error) {
	statement, err := r.connection.Prepare("SELECT * FROM prices WHERE world = $1 AND good_name = $2;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		id         string
		good_name  string
		world      string
		buy_price  int
		sell_price int
		created_at string
	)

	rows, err := statement.Query(worldName, good)

	defer rows.Close()

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var results []*Price

	for rows.Next() {
		scanErr := rows.Scan(&id, &good_name, &world, &buy_price, &sell_price, &created_at)

		if scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		createdAt, timeParseErr := time.Parse(time.DateOnly, created_at)

		if timeParseErr != nil {
			return nil, eris.New(timeParseErr.Error())
		}

		results = append(results, &Price{
			Id:        id,
			GoodName:  good_name,
			World:     world,
			BuyPrice:  buy_price,
			SellPrice: sell_price,
			CreatedAt: createdAt,
		})
	}

	return results, nil
}
