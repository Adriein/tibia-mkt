package script

import (
	"encoding/csv"
	"fmt"
	"github.com/rotisserie/eris"
	"math/rand"
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

		europeanCreatedAtFormat, timeParseErr := time.Parse("02-01-2006", record[0])

		if timeParseErr != nil {
			return nil, eris.New(timeParseErr.Error())
		}

		createdAtString := europeanCreatedAtFormat.Format(time.DateOnly)

		createdAt, stringToTimeParseErr := time.Parse(time.DateOnly, createdAtString)

		if stringToTimeParseErr != nil {
			return nil, eris.New(stringToTimeParseErr.Error())
		}

		// Generate random hours (0 to 23)
		randomHours := rand.Intn(24) // [0, 24)

		// Generate random minutes (0 to 59)
		randomMinutes := rand.Intn(60) // [0, 60)

		// Generate random seconds (0 to 59)
		randomSeconds := rand.Intn(60) // [0, 60)

		// Add the random durations to the original time
		timestampCreatedAt := createdAt.Add(time.Duration(randomHours)*time.Hour +
			time.Duration(randomMinutes)*time.Minute +
			time.Duration(randomSeconds)*time.Second)

		result = append(result, &CsvRow{
			SellPrice: price,
			CreatedAt: timestampCreatedAt,
		})
	}

	return result, nil
}
