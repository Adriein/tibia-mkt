package script

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/rotisserie/eris"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type SecuraPricesCsvRepository interface {
	Get(good string) ([]*CsvRow, error)
}

type SecuraPricesJsonRepository interface {
	Get(good string) ([]*JsonObj, error)
}

type CsvSecuraPricesRepository struct{}

func NewCsvSecuraPricesRepository() *CsvSecuraPricesRepository {
	return &CsvSecuraPricesRepository{}
}

type JsonSecuraPricesRepository struct{}

func NewJsonSecuraPricesRepository() *JsonSecuraPricesRepository {
	return &JsonSecuraPricesRepository{}
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

func (c *JsonSecuraPricesRepository) Get(good string) ([]*JsonObj, error) {
	path, _ := os.Getwd()

	filePath := fmt.Sprintf("%s/data/api.tibiamarket.top.%s.json", path, good)

	file, openErr := os.Open(filePath)

	if openErr != nil {
		return nil, eris.New(openErr.Error())
	}

	defer file.Close()

	reader := json.NewDecoder(file)

	var objs [][]*RawJsonObj

	readAllErr := reader.Decode(&objs)

	if readAllErr != nil {
		return nil, eris.New(readAllErr.Error())
	}

	var result []*JsonObj

	for _, obj := range objs[0] {
		seconds := int64(obj.Time)

		nanosecondsFloat := (obj.Time - float64(seconds)) * 1e9

		nanoseconds := int64(nanosecondsFloat)

		dateTime := time.Unix(seconds, nanoseconds)

		result = append(result, &JsonObj{
			Time:       dateTime,
			BuyOffer:   obj.BuyOffer,
			SellOffer:  obj.SellOffer,
			BuyOffers:  obj.BuyOffers,
			SellOffers: obj.SellOffers,
		})

		/*if timeParseErr != nil {
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
		})*/
	}

	return result, nil
}
