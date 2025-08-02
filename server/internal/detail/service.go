package detail

import (
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/statistics"
)

type Service struct {
	stats        *statistics.Statistics
	priceService *price.Service
}

func NewService(stats *statistics.Statistics, priceService *price.Service) *Service {
	return &Service{
		stats:        stats,
		priceService: priceService,
	}
}

func (s *Service) GetDetail(world string, good string) (*Detail, error) {
	prices, getPricesErr := s.priceService.GetPrices(world, good)

	if getPricesErr != nil {
		return nil, getPricesErr
	}

	var (
		sellPrices []int
		buyPrices  []int
	)

	for _, p := range prices {
		if p.OfferType == constants.SellOffer {
			sellPrices = append(sellPrices, p.UnitPrice)

			continue
		}

		buyPrices = append(buyPrices, p.UnitPrice)
	}

	buyOfferMean := s.stats.Mean(buyPrices)
	sellOfferMean := s.stats.Mean(sellPrices)

	buyOfferStdDeviation := s.stats.StdDeviation(buyPrices)
	sellOfferStdDeviation := s.stats.StdDeviation(sellPrices)

	buyOfferMedian := s.stats.Median(buyPrices)
	sellOfferMedian := s.stats.Median(sellPrices)

	return &Detail{
		SellOffersMean:         int(sellOfferMean),
		SellOffersStdDeviation: int(sellOfferStdDeviation),
		SellOffersMedian:       sellOfferMedian,
		BuyOffersMean:          int(buyOfferMean),
		BuyOffersStdDeviation:  int(buyOfferStdDeviation),
		BuyOffersMedian:        buyOfferMedian,
	}, nil
}
