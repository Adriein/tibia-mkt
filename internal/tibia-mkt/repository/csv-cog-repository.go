package repository

import (
	"encoding/csv"
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/types"
	"os"
	"strconv"
	"time"
)

type CsvSecuraCogRepository struct {
}

func NewCsvSecuraCogRepository() *CsvSecuraCogRepository {
	return &CsvSecuraCogRepository{}
}

func (c *CsvSecuraCogRepository) Find(criteria types.Criteria) ([]types.CogSku, error) {
	var result []types.CogSku

	path, _ := os.Getwd()

	file, openErr := os.Open(fmt.Sprintf("%s/data/Secura COG's - TC.csv", path))

	if openErr != nil {
		return nil, types.ApiError{
			Msg:      openErr.Error(),
			Function: "Find -> os.Open()",
			File:     "csv-cog-repository.go",
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

		result = append(result, types.CogSku{
			Date:  date,
			Price: price,
		})
	}

	return result, nil
}
