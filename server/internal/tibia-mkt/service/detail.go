package service

import (
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"sort"
	"time"
)

type KeyValue struct {
	Key   string
	Value int
}

type DetailService struct {
	goodRepository          types.Repository[types.Good]
	killStatisticRepository types.Repository[types.KillStatistic]
	dataSnapshotRepository  types.Repository[types.DataSnapshot]
	repoFactory             *helper.RepositoryFactory
	prob                    *helper.ProbHelper
}

func NewDetailService(
	goodRepository types.Repository[types.Good],
	killStatisticRepository types.Repository[types.KillStatistic],
	dataSnapshotRepository types.Repository[types.DataSnapshot],
	repoFactory *helper.RepositoryFactory,
	prob *helper.ProbHelper,
) *DetailService {
	return &DetailService{
		goodRepository:          goodRepository,
		killStatisticRepository: killStatisticRepository,
		dataSnapshotRepository:  dataSnapshotRepository,
		repoFactory:             repoFactory,
		prob:                    prob,
	}
}

func (s *DetailService) Execute(good string) (types.Detail, error) {
	goodDetail, goodErr := s.getGoodInformation(good)

	if goodErr != nil {
		return types.Detail{}, goodErr
	}

	goods, goodsDataErr := s.getGoodData(good)

	if goodsDataErr != nil {
		return types.Detail{}, goodsDataErr
	}

	killStatistics, killStatisticErr := s.getKillStatistics(goodDetail)

	if killStatisticErr != nil {
		killStatistics = make([]types.CreatureKillStatistic, 0)
	}

	frequencyChart, prices := s.buildPriceFrequencyChart(goods)

	historicData, historicDataErr := s.get15DaysSellOfferHistoricData(good)

	if historicDataErr != nil {
		return types.Detail{}, historicDataErr
	}

	if len(historicData) == 0 {
		historicData = make([]types.DataSnapshot, 0)
	}

	mean := s.computeMean(prices)

	return types.Detail{
		Wiki:                  goodDetail.Link,
		GoodRecord:            goods,
		Creatures:             killStatistics,
		SellPriceMean:         mean,
		StdDeviation:          s.prob.StdDeviation(prices),
		SellOfferFrequency:    frequencyChart,
		SellOfferHistoricData: historicData,
	}, nil
}

func (s *DetailService) getGoodInformation(itemName string) (types.Good, error) {
	var filters []types.Filter

	filters = append(filters, types.Filter{
		Name:    "name",
		Operand: constants.Equal,
		Value:   itemName,
	})

	criteria := types.Criteria{Filters: filters}

	result, err := s.goodRepository.FindOne(criteria)

	if err != nil {
		return types.Good{}, err
	}

	return result, nil
}

func (s *DetailService) getGoodData(good string) ([]types.GoodRecord, error) {
	repository := s.repoFactory.Get(good)

	var filters []types.Filter

	filters = append(filters, types.Filter{Name: "world", Operand: constants.Equal, Value: constants.WorldSecura})

	goods, repositoryErr := repository.Find(types.Criteria{Filters: filters})

	if repositoryErr != nil {
		return nil, repositoryErr
	}

	return goods, nil
}

func (s *DetailService) getKillStatistics(good types.Good) ([]types.CreatureKillStatistic, error) {
	var creatureKillStatistics []types.CreatureKillStatistic

	for _, creature := range good.Drop {
		var filters []types.Filter

		filters = append(
			filters,
			types.Filter{
				Name:    "creature_name",
				Operand: constants.Equal,
				Value:   creature.Name,
			},
			types.Filter{
				Name:    "created_at",
				Operand: constants.GreaterThanOrEqual,
				Value:   time.Now().Format(time.DateOnly),
			},
		)

		criteria := types.Criteria{Filters: filters}

		result, err := s.killStatisticRepository.FindOne(criteria)

		if err != nil {
			return nil, err
		}

		creatureKillStatistics = append(creatureKillStatistics, types.CreatureKillStatistic{
			Name:          result.CreatureName,
			DropRate:      result.DropRate,
			KillStatistic: result.AmountKilled,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
		})
	}

	return creatureKillStatistics, nil
}

func (s *DetailService) buildPriceFrequencyChart(goods []types.GoodRecord) ([]types.SellOfferFrequency, []int) {
	var frequencyResults []types.SellOfferFrequency

	priceRanges := make(map[string]int)
	prices := make([]int, len(goods))

	for i := 0; i < len(goods); i++ {
		prices[i] = goods[i].SellPrice

		rangeStart := (goods[i].SellPrice / 100) * 100
		rangeEnd := rangeStart + 99

		rangeStr := fmt.Sprintf("%d-%d", rangeStart, rangeEnd)

		priceRanges[rangeStr]++
	}

	sortedKeyValue := s.sortFrequencyMap(priceRanges)

	for _, keyValue := range sortedKeyValue {
		frequency := float64(keyValue.Value) / float64(len(goods))

		frequencyResults = append(
			frequencyResults,
			types.SellOfferFrequency{Range: keyValue.Key, Occurrences: keyValue.Value, Frequency: frequency},
		)
	}

	return frequencyResults, prices
}

func (s *DetailService) sortFrequencyMap(priceRanges map[string]int) []KeyValue {
	var result []KeyValue
	for k, v := range priceRanges {
		result = append(result, KeyValue{Key: k, Value: v})
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].Key < result[j].Key
	})

	return result
}

func (s *DetailService) get15DaysSellOfferHistoricData(good string) ([]types.DataSnapshot, error) {
	filters := make([]types.Filter, 3)

	now := time.Now()

	fifteenDaysAgo := now.AddDate(0, 0, -15).Format(time.DateOnly)

	filters[0] = types.Filter{Name: "cog", Operand: constants.Equal, Value: good}
	filters[1] = types.Filter{Name: "created_at", Operand: constants.GreaterThanOrEqual, Value: fifteenDaysAgo}
	filters[2] = types.Filter{Name: "created_at", Operand: constants.LessThanOrEqual, Value: now.Format(time.DateTime)}

	result, repoErr := s.dataSnapshotRepository.Find(types.Criteria{Filters: filters})

	if repoErr != nil {
		return nil, repoErr
	}

	return result, nil
}

func (s *DetailService) computeMean(prices []int) int {
	if len(prices) == 1 {
		return prices[0]
	}

	return int(s.prob.Mean(prices))
}
