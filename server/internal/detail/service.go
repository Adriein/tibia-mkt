package detail

import (
	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/statistics"
	"math"
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
		sellPrices          []int
		buyPrices           []int
		sellOfferMarketCap  int
		buyOfferMarketCap   int
		oneDayAgoMarketCap  int
		twoDaysAgoMarketCap int
		totalGoodsBeingSold int
	)

	oneDayAgo := time.Now().Add(time.Duration(-24) * time.Hour)
	twoDaysAgo := time.Now().Add(time.Duration(-48) * time.Hour)

	for _, p := range prices {
		if p.OfferType == constants.SellOffer {
			if p.EndAt.After(time.Now()) {
				sellOfferMarketCap += p.GoodAmount * p.UnitPrice
				totalGoodsBeingSold += p.GoodAmount
			}

			if p.CreatedAt.After(oneDayAgo) && p.CreatedAt.Before(time.Now()) {
				oneDayAgoMarketCap += p.GoodAmount * p.UnitPrice
			}

			if p.CreatedAt.After(twoDaysAgo) && p.CreatedAt.Before(oneDayAgo) {
				twoDaysAgoMarketCap += p.GoodAmount * p.UnitPrice
			}

			sellPrices = append(sellPrices, p.UnitPrice)

			continue
		}

		if p.EndAt.After(time.Now()) {
			buyOfferMarketCap += p.GoodAmount * p.UnitPrice
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

	stdDeviationRelativeToMean := (sellOfferStdDeviation / sellOfferMean) * 100

	marketCapDelta := float64(oneDayAgoMarketCap - twoDaysAgoMarketCap)

	marketVolumeTendencyPercentage := (marketCapDelta / float64(twoDaysAgoMarketCap)) * 100

	marketStatus := s.assertMarketStatus(stdDeviationRelativeToMean, spreadPercentage, marketVolumeTendencyPercentage)

	buyPressurePercentage := math.Round(float64(buyOfferMarketCap) / (float64(buyOfferMarketCap) + float64(sellOfferMarketCap)) * 100)
	sellPressurePercentage := math.Round(float64(sellOfferMarketCap) / (float64(buyOfferMarketCap) + float64(sellOfferMarketCap)) * 100)

	spreadScore := math.Max(0, 1-spreadPercentage/0.15)
	volumeScore := math.Max(float64(oneDayAgoMarketCap)/float64(sellOfferMarketCap), 1)

	liquidity := int((spreadScore + volumeScore) / 2 * 100)

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
			BuySellSpread:                  buySellSpread,
			SpreadPercentage:               int(spreadPercentage),
			MarketCap:                      sellOfferMarketCap,
			LastTwentyFourHoursVolume:      oneDayAgoMarketCap,
			MarketStatus:                   marketStatus,
			MarketVolumePercentageTendency: int(marketVolumeTendencyPercentage),
			TotalGoodsBeingSold:            totalGoodsBeingSold,
		},
		Insights: DetailInsights{
			MarketType:   "unknown",
			BuyPressure:  int(buyPressurePercentage),
			SellPressure: int(sellPressurePercentage),
			Liquidity:    liquidity,
		},
	}, nil
}

func (s *Service) assertMarketStatus(
	stdDeviationRelativeToMean float64,
	spreadPercentage float64,
	marketVolumeTendencyPercentage float64,
) string {
	score := 0

	if stdDeviationRelativeToMean >= 15 {
		score += 2
	}

	if stdDeviationRelativeToMean >= 5 {
		score += 1
	}

	if spreadPercentage >= 8 {
		score += 2
	}

	if spreadPercentage >= 3 {
		score += 1
	}

	if marketVolumeTendencyPercentage <= -50 {
		score += 2
	}

	if marketVolumeTendencyPercentage <= -30 {
		score += 1
	}

	if score >= 4 {
		return constants.VolatileMarketStatus
	}

	if score >= 2 {
		return constants.RiskyMarketStatus
	}

	return constants.StableMarketStatus
}
