package detail

import (
	"github.com/adriein/tibia-mkt/internal/price"
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

	sellPrices := make([]int, len(prices))
	buyPrices := make([]int, len(prices))

	for _, p := range prices {
		sellPrices = append(sellPrices, p.SellPrice)
		buyPrices = append(buyPrices, p.BuyPrice)
	}

	buyOfferMean := s.stats.Mean(buyPrices)
	sellOfferMean := s.stats.Mean(sellPrices)

	buyOfferStdDeviation := s.stats.StdDeviation(buyPrices)
	sellOfferStdDeviation := s.stats.StdDeviation(sellPrices)

	buyOfferMedian := s.stats.Median(buyPrices)
	sellOfferMedian := s.stats.Median(sellPrices)

	return &Detail{
		SellOffersMean:         int(sellOfferMean),
		SellOffersStdDeviation: sellOfferStdDeviation,
		SellOffersMedian:       sellOfferMedian,
		BuyOffersMean:          int(buyOfferMean),
		BuyOffersStdDeviation:  buyOfferStdDeviation,
		BuyOffersMedian:        buyOfferMedian,
	}, nil
}
