package price

import (
	"database/sql"
	"strings"
	"time"

	"github.com/rotisserie/eris"
)

type PriceRepository interface {
	FindNewestOfferByGoodAndWorld(world string, good string, offerType string) ([]*Price, error)
	Save(price *Price) error
}

type PgPriceRepository struct {
	connection *sql.DB
}

func NewPgPriceRepository(connection *sql.DB) *PgPriceRepository {
	return &PgPriceRepository{
		connection: connection,
	}
}

func (r *PgPriceRepository) FindNewestOfferByGoodAndWorld(worldName string, good string, offerType string) ([]*Price, error) {
	statement, err := r.connection.Prepare("SELECT * FROM prices WHERE world = $1 AND good_name = $2 AND offer_type = $3 ORDER BY created_at;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		id          string
		batch_id    int
		market_id   string
		offer_type  string
		good_name   string
		world       string
		created_by  string
		good_amount int
		unit_price  int
		total_price int
		end_at      string
		created_at  string
	)

	rows, err := statement.Query(worldName, good, offerType)

	defer rows.Close()

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var results []*Price

	for rows.Next() {
		scanErr := rows.Scan(&id, &batch_id, &market_id, &offer_type, &good_name, &world, &created_by, &good_amount, &unit_price, &total_price, &end_at, &created_at)

		if scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		createdAt, createdAtTimeParseErr := time.Parse(time.RFC3339Nano, created_at)

		if createdAtTimeParseErr != nil {
			return nil, eris.New(createdAtTimeParseErr.Error())
		}

		endAt, endAtTimeParseErr := time.Parse(time.RFC3339Nano, end_at)

		if endAtTimeParseErr != nil {
			return nil, eris.New(endAtTimeParseErr.Error())
		}

		results = append(results, &Price{
			Id:         id,
			BatchId:    batch_id,
			MarketId:   market_id,
			OfferType:  offer_type,
			GoodName:   good_name,
			World:      world,
			CreatedBy:  created_by,
			GoodAmount: good_amount,
			UnitPrice:  unit_price,
			TotalPrice: total_price,
			EndAt:      endAt,
			CreatedAt:  createdAt,
		})
	}

	return results, nil
}

func (r *PgPriceRepository) Save(price *Price) error {
	var b strings.Builder
	b.WriteString("INSERT INTO prices (")
	b.WriteString("id, batch_id, market_id, offer_type, good_name, world, created_by, good_amount, unit_price, ")
	b.WriteString("total_price, end_at, created_at")
	b.WriteString(") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)")

	var query = b.String()

	_, err := r.connection.Exec(
		query,
		price.Id,
		price.BatchId,
		price.MarketId,
		price.OfferType,
		price.GoodName,
		price.World,
		price.CreatedBy,
		price.GoodAmount,
		price.UnitPrice,
		price.TotalPrice,
		price.EndAt.Format(time.DateTime),
		price.CreatedAt.Format(time.DateTime),
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}
