package service

import (
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"sort"
	"time"
)

type KeyValue struct {
	Key   string
	Value int
}

type DetailService struct {
	cogRepository           types.Repository[types.Cog]
	killStatisticRepository types.Repository[types.KillStatistic]
	dataSnapshotRepository  types.Repository[types.DataSnapshot]
	repoFactory             *RepositoryFactory
	prob                    *ProbHelper
}

func NewDetailService(
	cogRepository types.Repository[types.Cog],
	killStatisticRepository types.Repository[types.KillStatistic],
	dataSnapshotRepository types.Repository[types.DataSnapshot],
	repoFactory *RepositoryFactory,
	prob *ProbHelper,
) *DetailService {
	return &DetailService{
		cogRepository:           cogRepository,
		killStatisticRepository: killStatisticRepository,
		dataSnapshotRepository:  dataSnapshotRepository,
		repoFactory:             repoFactory,
		prob:                    prob,
	}
}

func (s *DetailService) Execute(cog string) (types.Detail, error) {
	cogDetail, cogErr := s.getCogInformation(cog)

	if cogErr != nil {
		return types.Detail{}, cogErr
	}

	cogs, cogDataErr := s.getCogData(cog)

	if cogDataErr != nil {
		return types.Detail{}, cogDataErr
	}

	killStatistics, killStatisticErr := s.getKillStatistics(cogDetail)

	if killStatisticErr != nil {
		killStatistics = make([]types.CreatureKillStatistic, 0)
	}

	frequencyChart, prices := s.buildPriceFrequencyChart(cogs)

	historicData, historicDataErr := s.get15DaysSellOfferHistoricData(cog)

	if historicDataErr != nil {
		return types.Detail{}, historicDataErr
	}

	if len(historicData) == 0 {
		historicData = make([]types.DataSnapshot, 0)
	}

	return types.Detail{
		Wiki:                  cogDetail.Link,
		Cog:                   cogs,
		Creatures:             killStatistics,
		SellPriceMean:         int(s.prob.Mean(prices)),
		StdDeviation:          s.prob.StdDeviation(prices),
		SellOfferFrequency:    frequencyChart,
		SellOfferHistoricData: historicData,
	}, nil
}

func (s *DetailService) getCogInformation(itemName string) (types.Cog, error) {
	var filters []types.Filter

	filters = append(filters, types.Filter{
		Name:    "name",
		Operand: constants.Equal,
		Value:   itemName,
	})

	criteria := types.Criteria{Filters: filters}

	result, err := s.cogRepository.FindOne(criteria)

	if err != nil {
		return types.Cog{}, err
	}

	return result, nil
}

func (s *DetailService) getCogData(cog string) ([]types.CogSku, error) {
	repository := s.repoFactory.Get(cog)

	var filters []types.Filter

	filters = append(filters, types.Filter{Name: "world", Operand: constants.Equal, Value: constants.WorldSecura})

	cogs, repositoryErr := repository.Find(types.Criteria{Filters: filters})

	if repositoryErr != nil {
		return nil, repositoryErr
	}

	return cogs, nil
}

func (s *DetailService) getKillStatistics(cog types.Cog) ([]types.CreatureKillStatistic, error) {
	var creatureKillStatistics []types.CreatureKillStatistic

	for _, creature := range cog.Creatures {
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

func (s *DetailService) buildPriceFrequencyChart(cogs []types.CogSku) ([]types.SellOfferFrequency, []int) {
	var frequencyResults []types.SellOfferFrequency

	priceRanges := make(map[string]int)
	prices := make([]int, len(cogs))

	for i := 0; i < len(cogs); i++ {
		prices[i] = cogs[i].SellPrice

		rangeStart := (cogs[i].SellPrice / 100) * 100
		rangeEnd := rangeStart + 99

		rangeStr := fmt.Sprintf("%d-%d", rangeStart, rangeEnd)

		priceRanges[rangeStr]++
	}

	sortedKeyValue := s.sortFrequencyMap(priceRanges)

	for _, keyValue := range sortedKeyValue {
		frequency := float64(keyValue.Value) / float64(len(cogs))

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

func (s *DetailService) get15DaysSellOfferHistoricData(cog string) ([]types.DataSnapshot, error) {
	filters := make([]types.Filter, 3)

	now := time.Now()

	fifteenDaysAgo := now.AddDate(0, 0, -15).Format(time.DateOnly)

	filters[0] = types.Filter{Name: "cog", Operand: constants.Equal, Value: cog}
	filters[1] = types.Filter{Name: "created_at", Operand: constants.GreaterThanOrEqual, Value: fifteenDaysAgo}
	filters[2] = types.Filter{Name: "created_at", Operand: constants.LessThanOrEqual, Value: now.Format(time.DateTime)}

	result, repoErr := s.dataSnapshotRepository.Find(types.Criteria{Filters: filters})

	if repoErr != nil {
		return nil, repoErr
	}

	return result, nil
}
