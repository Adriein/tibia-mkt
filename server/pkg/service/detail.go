package service

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type DetailService struct {
	cogRepository           types.Repository[types.Cog]
	killStatisticRepository types.Repository[types.KillStatistic]
	repoFactory             *RepositoryFactory
}

func NewDetailService(
	cogRepository types.Repository[types.Cog],
	killStatisticRepository types.Repository[types.KillStatistic],
	repoFactory *RepositoryFactory,
) *DetailService {
	return &DetailService{
		cogRepository:           cogRepository,
		killStatisticRepository: killStatisticRepository,
		repoFactory:             repoFactory,
	}
}

func (s *DetailService) Execute(cog string) (types.DetailPresenterInput, error) {
	cogDetail, cogErr := s.getCogInformation(cog)

	if cogErr != nil {
		return types.DetailPresenterInput{}, cogErr
	}

	repository := s.repoFactory.Get(cog)

	var filters []types.Filter

	filters = append(filters, types.Filter{Name: "world", Operand: constants.Equal, Value: "Secura"})

	results, repositoryErr := repository.Find(types.Criteria{Filters: filters})

	if repositoryErr != nil {
		return types.DetailPresenterInput{}, repositoryErr
	}

	killStatistics, killStatisticErr := s.getKillStatistics(cogDetail)

	if killStatisticErr != nil {
		return types.DetailPresenterInput{}, killStatisticErr
	}

	return types.DetailPresenterInput{
		Wiki:      cogDetail.Link,
		Cog:       results,
		Creatures: killStatistics,
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
		})
	}

	return creatureKillStatistics, nil
}
