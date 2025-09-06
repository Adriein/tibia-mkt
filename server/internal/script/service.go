package script

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/adriein/tibia-mkt/internal/event"
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/google/uuid"
)

type Service struct {
	csvDataRepository  SecuraPricesCsvRepository
	jsonDataRepository SecuraPricesJsonRepository
	pricesRepository   price.PriceRepository
	eventsRepository   event.EventRepository
}

func NewService(
	csvDataRepository SecuraPricesCsvRepository,
	jsonDataRepository SecuraPricesJsonRepository,
	pricesRepository price.PriceRepository,
	eventsRepository event.EventRepository,
) *Service {
	return &Service{
		csvDataRepository:  csvDataRepository,
		jsonDataRepository: jsonDataRepository,
		pricesRepository:   pricesRepository,
		eventsRepository:   eventsRepository,
	}
}

func (s *Service) SeedPricesFromCsv() error {
	goods := [6]string{
		constants.HoneycombEntity,
		constants.SwamplingWoodEntity,
		constants.TibiaCoinEntity,
		constants.BrokenShamanicStaffEntity,
		constants.TurtleShell,
		constants.CobraRod,
	}

	for _, good := range goods {
		csvRow, err := s.csvDataRepository.Get(good)

		if err != nil {
			return err
		}

		for _, row := range csvRow {
			var sb strings.Builder
			id := uuid.New().String()
			endAt := row.CreatedAt.AddDate(0, 0, 30)

			sb.WriteString(fmt.Sprintf("%d", endAt.Unix()))
			sb.WriteString(constants.SellOffer)
			sb.WriteString(constants.WorldSecura)
			sb.WriteString(strconv.Itoa(row.SellPrice))

			marketId := sb.String()

			sb.Reset()

			randomSellAmount := rand.Intn(100) + 1

			sellRegisteredPrice := &price.Price{
				Id:         id,
				BatchId:    1,
				MarketId:   marketId,
				OfferType:  constants.SellOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: randomSellAmount,
				UnitPrice:  row.SellPrice,
				TotalPrice: row.SellPrice * randomSellAmount,
				EndAt:      endAt,
				CreatedAt:  row.CreatedAt,
			}

			if saveErr := s.pricesRepository.Save(sellRegisteredPrice); saveErr != nil {
				return saveErr
			}

			id = uuid.New().String()

			randomBuyPrice := rand.Intn(3000) + 1
			randomBuyAmount := rand.Intn(100) + 1

			unitPrice := row.SellPrice - randomBuyPrice

			sb.WriteString(fmt.Sprintf("%d", endAt.Unix()))
			sb.WriteString(constants.BuyOffer)
			sb.WriteString(constants.WorldSecura)
			sb.WriteString(strconv.Itoa(unitPrice))

			marketId = sb.String()

			buyRegisteredPrice := &price.Price{
				Id:         id,
				OfferType:  constants.BuyOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: randomBuyAmount,
				UnitPrice:  unitPrice,
				TotalPrice: row.SellPrice * randomBuyAmount,
				EndAt:      endAt,
				CreatedAt:  row.CreatedAt,
			}

			if saveErr := s.pricesRepository.Save(buyRegisteredPrice); saveErr != nil {
				return saveErr
			}
		}
	}

	return nil
}

func (s *Service) SeedPricesFromExternalApiJson() error {
	goods := [3]string{
		constants.TibiaCoinEntity,
		constants.HoneycombEntity,
		constants.SwamplingWoodEntity,
	}

	for _, good := range goods {
		jsonObjs, err := s.jsonDataRepository.Get(good)

		if err != nil {
			return err
		}

		batchId := 1

		for _, row := range jsonObjs {
			var sb strings.Builder
			id := uuid.New().String()
			endAt := row.Time.AddDate(0, 0, 30)

			sb.WriteString(fmt.Sprintf("%d", endAt.Unix()))
			sb.WriteString(constants.SellOffer)
			sb.WriteString(constants.WorldSecura)
			sb.WriteString(strconv.Itoa(row.SellOffer))

			marketId := sb.String()

			sb.Reset()

			sellRegisteredPrice := &price.Price{
				Id:         id,
				BatchId:    batchId,
				MarketId:   marketId,
				OfferType:  constants.SellOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: row.SellOffers,
				UnitPrice:  row.SellOffer,
				TotalPrice: row.SellOffer * row.SellOffers,
				EndAt:      endAt,
				CreatedAt:  row.Time,
			}

			if saveErr := s.pricesRepository.Save(sellRegisteredPrice); saveErr != nil {
				return saveErr
			}

			id = uuid.New().String()

			sb.WriteString(fmt.Sprintf("%d", endAt.Unix()))
			sb.WriteString(constants.BuyOffer)
			sb.WriteString(constants.WorldSecura)
			sb.WriteString(strconv.Itoa(row.BuyOffer))

			marketId = sb.String()

			buyRegisteredPrice := &price.Price{
				Id:         id,
				BatchId:    batchId,
				MarketId:   marketId,
				OfferType:  constants.BuyOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: row.BuyOffers,
				UnitPrice:  row.BuyOffer,
				TotalPrice: row.BuyOffer * row.BuyOffers,
				EndAt:      endAt,
				CreatedAt:  row.Time,
			}

			if saveErr := s.pricesRepository.Save(buyRegisteredPrice); saveErr != nil {
				return saveErr
			}

			batchId++
		}

		priceRegisteredEvent := &event.Event{
			Name:        constants.EventDataIngestion,
			GoodName:    good,
			World:       constants.WorldSecura,
			Description: constants.EventDataIngestionDescription,
			OccurredAt:  time.Now(),
		}

		if saveEventErr := s.eventsRepository.Save(priceRegisteredEvent); saveEventErr != nil {
			return saveEventErr
		}
	}

	return nil
}
