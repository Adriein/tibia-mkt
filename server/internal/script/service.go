package script

import (
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/google/uuid"
	"math/rand"
)

type Service struct {
	rawDataRepository SecuraPricesRepository
	pricesRepository  price.PriceRepository
}

func NewService(
	rawDataRepository SecuraPricesRepository,
	pricesRepository price.PriceRepository,
) *Service {
	return &Service{
		rawDataRepository: rawDataRepository,
		pricesRepository:  pricesRepository,
	}
}

func (s *Service) SeedPrices() error {
	goods := [4]string{
		constants.HoneycombEntity,
		constants.SwamplingWoodEntity,
		constants.TibiaCoinEntity,
		constants.BrokenShamanicStaffEntity,
	}

	for _, good := range goods {
		csvRow, err := s.rawDataRepository.Get(good)

		if err != nil {
			return err
		}

		for _, row := range csvRow {
			id := uuid.New().String()

			randomSellAmount := rand.Intn(100) + 1

			sellRegisteredPrice := &price.Price{
				Id:         id,
				OfferType:  constants.SellOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: randomSellAmount,
				UnitPrice:  row.SellPrice,
				TotalPrice: row.SellPrice * randomSellAmount,
				EndAt:      row.CreatedAt.AddDate(0, 0, 30),
				CreatedAt:  row.CreatedAt,
			}

			if saveErr := s.pricesRepository.Save(sellRegisteredPrice); saveErr != nil {
				return saveErr
			}

			id = uuid.New().String()

			randomBuyPrice := rand.Intn(3000) + 1
			randomBuyAmount := rand.Intn(100) + 1

			buyRegisteredPrice := &price.Price{
				Id:         id,
				OfferType:  constants.BuyOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: randomBuyAmount,
				UnitPrice:  row.SellPrice - randomBuyPrice,
				TotalPrice: row.SellPrice * randomBuyAmount,
				EndAt:      row.CreatedAt.AddDate(0, 0, 30),
				CreatedAt:  row.CreatedAt,
			}

			if saveErr := s.pricesRepository.Save(buyRegisteredPrice); saveErr != nil {
				return saveErr
			}
		}
	}

	return nil
}
