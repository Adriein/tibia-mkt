package script

import (
	"encoding/csv"
	"fmt"
	"github.com/rotisserie/eris"
	"os"
	"strconv"
	"time"
)

type SecuraPricesRepository interface {
	Get(good string) ([]*CsvRow, error)
}

type CsvSecuraPricesRepository struct{}

func NewCsvSecuraPricesRepository() *CsvSecuraPricesRepository {
	return &CsvSecuraPricesRepository{}
}

func (c *CsvSecuraPricesRepository) Get(good string) ([]*CsvRow, error) {
	path, _ := os.Getwd()

	filePath := fmt.Sprintf("%s/data/%s-cog.csv", path, good)

	file, openErr := os.Open(filePath)

	if openErr != nil {
		return nil, eris.New(openErr.Error())
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, readAllErr := reader.ReadAll()

	if readAllErr != nil {
		return nil, eris.New(readAllErr.Error())
	}

	var result []*CsvRow

	for _, record := range records {
		price, intParseErr := strconv.Atoi(record[1])

		if intParseErr != nil {
			return nil, eris.New(intParseErr.Error())
		}

		createdAt, timeParseErr := time.Parse(time.DateOnly, record[0])

		if timeParseErr != nil {
			return nil, eris.New(timeParseErr.Error())
		}

		result = append(result, &CsvRow{
			SellPrice: price,
			CreatedAt: createdAt,
		})
	}

	return result, nil
}
