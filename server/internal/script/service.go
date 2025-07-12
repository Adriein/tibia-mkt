package script

import (
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/google/uuid"
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

			registeredPrice := &price.Price{
				Id:        id,
				GoodName:  good,
				World:     constants.WorldSecura,
				BuyPrice:  row.SellPrice - 1000,
				SellPrice: row.SellPrice,
				CreatedAt: row.CreatedAt,
			}

			if saveErr := s.pricesRepository.Save(registeredPrice); saveErr != nil {
				return saveErr
			}
		}
	}

	return nil
}
