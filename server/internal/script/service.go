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

			randomAmount := rand.Intn(100) + 1

			registeredPrice := &price.Price{
				Id:         id,
				OfferType:  constants.SellOffer,
				GoodName:   good,
				World:      constants.WorldSecura,
				CreatedBy:  "anonymous",
				GoodAmount: randomAmount,
				UnitPrice:  row.SellPrice,
				TotalPrice: row.SellPrice * randomAmount,
				EndAt:      row.CreatedAt.AddDate(0, 0, 30),
				CreatedAt:  row.CreatedAt,
			}

			if saveErr := s.pricesRepository.Save(registeredPrice); saveErr != nil {
				return saveErr
			}
		}
	}

	return nil
}
