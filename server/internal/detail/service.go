package detail

import (
	"fmt"
	"math"
	"time"

	"github.com/adriein/tibia-mkt/internal/price"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/statistics"
	"github.com/rotisserie/eris"
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

func (s *Service) GetDetail(world, good string) (*Detail, error) {
	now := time.Now()

	prices, err := s.priceService.GetPrices(world, good)

	if err != nil {
		return nil, err
	}

	metrics := s.aggregatePrices(prices, now)

	if len(metrics.sellPrices) == 0 || len(metrics.buyPrices) == 0 {
		return nil, eris.New(fmt.Sprintf("insufficient market data for %s in %s", good, world))
	}

	stats := s.calculateStats(metrics)
	insights := s.calculateInsights(metrics, stats)
	overview := s.buildOverview(metrics)

	return &Detail{
		Stats:    stats,
		Overview: overview,
		Insights: insights,
	}, nil
}

type marketMetrics struct {
	sellPrices, buyPrices       []int
	sellOfferMarketCap          int
	buyOfferMarketCap           int
	oneDayAgoMarketCap          int
	twoDaysAgoMarketCap         int
	totalGoodsBeingSold         int
	mostRecentSellOffer         int
	mostRecentBuyOffer          int
	marketVolumeTendencyPercent float64
}

func (s *Service) aggregatePrices(prices []*price.Price, now time.Time) marketMetrics {
	var m marketMetrics

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-48 * time.Hour)

	for _, p := range prices {
		if p.UnitPrice == -1 {
			continue
		}

		if p.OfferType == constants.SellOffer {
			if p.EndAt.After(now) {
				m.sellOfferMarketCap += p.GoodAmount * p.UnitPrice
				m.totalGoodsBeingSold += p.GoodAmount
			}

			if p.CreatedAt.After(oneDayAgo) && p.CreatedAt.Before(now) {
				m.oneDayAgoMarketCap += p.GoodAmount * p.UnitPrice
			}

			if p.CreatedAt.After(twoDaysAgo) && p.CreatedAt.Before(oneDayAgo) {
				m.twoDaysAgoMarketCap += p.GoodAmount * p.UnitPrice
			}

			m.sellPrices = append(m.sellPrices, p.UnitPrice)

			continue
		}

		if p.EndAt.After(now) {
			m.buyOfferMarketCap += p.GoodAmount * p.UnitPrice
		}

		m.buyPrices = append(m.buyPrices, p.UnitPrice)
	}

	if len(m.sellPrices) > 0 {
		m.mostRecentSellOffer = m.sellPrices[len(m.sellPrices)-1]
	}

	if len(m.buyPrices) > 0 {
		m.mostRecentBuyOffer = m.buyPrices[len(m.buyPrices)-1]
	}

	if m.twoDaysAgoMarketCap > 0 {
		delta := float64(m.oneDayAgoMarketCap - m.twoDaysAgoMarketCap)

		m.marketVolumeTendencyPercent = (delta / float64(m.twoDaysAgoMarketCap)) * 100
	}

	return m
}

func (s *Service) calculateStats(m marketMetrics) DetailStats {
	return DetailStats{
		SellOffersMean:         int(s.stats.Mean(m.sellPrices)),
		SellOffersStdDeviation: int(s.stats.StdDeviation(m.sellPrices)),
		SellOffersMedian:       s.stats.Median(m.sellPrices),
		BuyOffersMean:          int(s.stats.Mean(m.buyPrices)),
		BuyOffersStdDeviation:  int(s.stats.StdDeviation(m.buyPrices)),
		BuyOffersMedian:        s.stats.Median(m.buyPrices),
	}
}

func (s *Service) calculateInsights(m marketMetrics, stats DetailStats) DetailInsights {
	buySellSpread := m.mostRecentSellOffer - m.mostRecentBuyOffer

	spreadPercentage := helper.PercentSafe(buySellSpread, m.mostRecentSellOffer)

	stdDeviationRelativeToMean := helper.PercentSafe(
		float64(stats.SellOffersStdDeviation),
		float64(stats.SellOffersMean),
	)

	marketStatus := s.assertMarketStatus(stdDeviationRelativeToMean, spreadPercentage, m.marketVolumeTendencyPercent)

	buyPressure := helper.PercentSafe(float64(m.buyOfferMarketCap), float64(m.buyOfferMarketCap+m.sellOfferMarketCap))
	sellPressure := helper.PercentSafe(float64(m.sellOfferMarketCap), float64(m.buyOfferMarketCap+m.sellOfferMarketCap))

	spreadScore := math.Max(0, 1-spreadPercentage/0.15)

	volumeScore := 1.0

	if m.sellOfferMarketCap > 0 {
		volumeScore = math.Max(float64(m.oneDayAgoMarketCap)/float64(m.sellOfferMarketCap), 1)
	}

	liquidity := int((spreadScore + volumeScore) / 2 * 100)

	marketType := s.assertMarketTendency(m.marketVolumeTendencyPercent, buyPressure)

	return DetailInsights{
		MarketStatus: marketStatus,
		MarketType:   marketType,
		BuyPressure:  int(math.Round(buyPressure)),
		SellPressure: int(math.Round(sellPressure)),
		Liquidity:    liquidity,
	}
}

func (s *Service) buildOverview(m marketMetrics) DetailOverview {
	buySellSpread := m.mostRecentSellOffer - m.mostRecentBuyOffer
	spreadPercentage := int(helper.PercentSafe(buySellSpread, m.mostRecentSellOffer))

	return DetailOverview{
		BuySellSpread:                  buySellSpread,
		SpreadPercentage:               spreadPercentage,
		MarketCap:                      m.sellOfferMarketCap,
		LastTwentyFourHoursVolume:      m.oneDayAgoMarketCap,
		MarketVolumePercentageTendency: int(m.marketVolumeTendencyPercent),
		TotalGoodsBeingSold:            m.totalGoodsBeingSold,
	}
}

func (s *Service) assertMarketTendency(marketVolumeTendencyPercentage float64, buyPressurePercentage float64) string {
	if marketVolumeTendencyPercentage < 10 && buyPressurePercentage > 55 {
		return constants.BullMarketType
	}
	if marketVolumeTendencyPercentage < -10 && buyPressurePercentage < 45 {
		return constants.BearMarketType
	}
	if math.Abs(marketVolumeTendencyPercentage) < 0.1 {
		return constants.SidewaysMarketType
	}
	if marketVolumeTendencyPercentage > 10 && buyPressurePercentage < 45 {
		return constants.BullExhaustionMarketType
	}
	if marketVolumeTendencyPercentage < -10 && buyPressurePercentage > 55 {
		return constants.PullbackMarketType
	}
	return constants.UnclearMarketType
}

func (s *Service) assertMarketStatus(stdDeviationRelativeToMean, spreadPercentage, marketVolumeTendencyPercentage float64) string {
	score := 0

	if stdDeviationRelativeToMean >= 15 {
		score += 2
	} else if stdDeviationRelativeToMean >= 5 {
		score++
	}

	if spreadPercentage >= 8 {
		score += 2
	} else if spreadPercentage >= 3 {
		score++
	}

	if marketVolumeTendencyPercentage <= -50 {
		score += 2
	} else if marketVolumeTendencyPercentage <= -30 {
		score++
	}

	if score >= 4 {
		return constants.VolatileMarketStatus
	}
	if score >= 2 {
		return constants.RiskyMarketStatus
	}
	return constants.StableMarketStatus
}
