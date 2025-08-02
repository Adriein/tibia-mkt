package detail

import (
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/statistics"
	"time"
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
		sellPrices              []int
		buyPrices               []int
		marketCap               int
		twentyFourHourMarketCap int
	)

	twentyFourHoursAgo := time.Now().Add(time.Duration(-24) * time.Hour)

	for _, p := range prices {
		if p.OfferType == constants.SellOffer {
			if p.EndAt.After(time.Now()) {
				marketCap += p.GoodAmount * p.UnitPrice
			}

			if p.EndAt.After(twentyFourHoursAgo) && p.EndAt.Before(time.Now()) {
				twentyFourHourMarketCap += p.GoodAmount * p.UnitPrice
			}

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

	mostRecentSellOffer := sellPrices[len(sellPrices)-1]
	mostRecentBuyOffer := buyPrices[len(buyPrices)-1]

	buySellSpread := mostRecentSellOffer - mostRecentBuyOffer
	spreadPercentage := (float64(buySellSpread) / float64(mostRecentSellOffer)) * 100

	stdDeviationRelativeToMean := (float64(sellOfferStdDeviation) / float64(sellOfferMean)) * 100

	marketStatus := s.assertMarketStatus(stdDeviationRelativeToMean, spreadPercentage)

	return &Detail{
		Stats: DetailStats{
			SellOffersMean:         int(sellOfferMean),
			SellOffersStdDeviation: int(sellOfferStdDeviation),
			SellOffersMedian:       sellOfferMedian,
			BuyOffersMean:          int(buyOfferMean),
			BuyOffersStdDeviation:  int(buyOfferStdDeviation),
			BuyOffersMedian:        buyOfferMedian,
		},
		Overview: DetailOverview{
			BuySellSpread:             buySellSpread,
			SpreadPercentage:          int(spreadPercentage),
			MarketCap:                 marketCap,
			LastTwentyFourHoursVolume: twentyFourHourMarketCap,
			MarketStatus:              marketStatus,
		},
	}, nil
}

func (s *Service) assertMarketStatus(stdDeviationRelativeToMean float64, spreadPercentage float64) string {
	if stdDeviationRelativeToMean >= 15 && spreadPercentage >= 8 {
		return constants.VolatileMarketStatus
	}

	if stdDeviationRelativeToMean >= 5 && spreadPercentage >= 3 {
		return constants.RiskyMarketStatus
	}

	return constants.StableMarketStatus
}
