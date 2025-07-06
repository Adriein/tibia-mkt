package repository

import (
	"encoding/csv"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type CsvSecuraCogRepository struct {
	name string
}

func NewCsvSecuraCogRepository() *CsvSecuraCogRepository {
	return &CsvSecuraCogRepository{
		name: "none",
	}
}

func (c *CsvSecuraCogRepository) Find(criteria types.Criteria) ([]types.GoodRecord, error) {
	var result []types.GoodRecord

	path, _ := os.Getwd()

	filePath := fmt.Sprintf("%s/data/%s-cog.csv", path, criteria.Filters[0].Value)

	file, openErr := os.Open(filePath)

	if openErr != nil {
		return nil, types.ApiError{
			Msg:      openErr.Error(),
			Function: "Find -> os.Open()",
			File:     "csv-cog-repository.go",
			Values:   []string{filePath},
		}
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, readAllErr := reader.ReadAll()

	if readAllErr != nil {
		return nil, types.ApiError{
			Msg:      readAllErr.Error(),
			Function: "Find -> reader.ReadAll()",
			File:     "csv-cog-repository.go",
		}
	}

	for _, record := range records {
		price, intParseErr := strconv.Atoi(record[1])

		if intParseErr != nil {
			return nil, types.ApiError{
				Msg:      intParseErr.Error(),
				Function: "Find -> strconv.Atoi()",
				File:     "csv-cog-repository.go",
			}
		}

		date, timeParseErr := time.Parse("02-01-2006", record[0])

		if timeParseErr != nil {
			return nil, types.ApiError{
				Msg:      timeParseErr.Error(),
				Function: "Find -> time.Parse()",
				File:     "csv-cog-repository.go",
			}
		}

		id := uuid.New()

		result = append(result, types.GoodRecord{
			Id:        id.String(),
			ItemName:  constants.TibiaCoinEntity,
			Date:      date,
			BuyPrice:  price - 1000,
			SellPrice: price,
			World:     constants.WorldSecura,
		})
	}

	return result, nil
}

func (c *CsvSecuraCogRepository) Save(entity types.GoodRecord) error {
	return nil
}

func (c *CsvSecuraCogRepository) GoodName() string {
	return c.name
}
