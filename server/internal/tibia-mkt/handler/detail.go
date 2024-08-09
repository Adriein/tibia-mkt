package handler

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
	"time"
)

type DetailHandler struct {
	cogRepository           types.Repository[types.Cog]
	killStatisticRepository types.Repository[types.KillStatistic]
	repoFactory             *service.RepositoryFactory
	presenter               types.Presenter
}

func NewDetailHandler(
	cogRepository types.Repository[types.Cog],
	killStatisticRepository types.Repository[types.KillStatistic],
	factory *service.RepositoryFactory,
	presenter types.Presenter,
) *DetailHandler {
	return &DetailHandler{
		cogRepository:           cogRepository,
		killStatisticRepository: killStatisticRepository,
		repoFactory:             factory,
		presenter:               presenter,
	}
}

type DetailHandlerPresenterInput struct {
	Wiki      string
	Cog       []types.CogSku
	Creatures []types.CreatureKillStatistic
}

func (h *DetailHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	paramsMap := r.URL.Query()

	if !paramsMap.Has("item") {
		return types.ApiError{
			Msg:      constants.NoCogSearchParamProvided,
			Function: "HomeHandler",
			File:     "home.go",
			Domain:   true,
		}
	}

	cog := paramsMap["item"][0]

	cogDetail, cogErr := h.getCogInformation(cog)

	if cogErr != nil {
		return cogErr
	}

	repository := h.repoFactory.Get(cog)

	var filters []types.Filter

	filters = append(filters, types.Filter{Name: "world", Operand: constants.Equal, Value: "Secura"})

	results, repositoryErr := repository.Find(types.Criteria{Filters: filters})

	if repositoryErr != nil {
		return repositoryErr
	}

	killStatistics, killStatisticErr := h.getKillStatistics(cogDetail)

	if killStatisticErr != nil {
		return killStatisticErr
	}

	response, presenterErr := h.presenter.Format(DetailHandlerPresenterInput{
		Wiki:      cogDetail.Link,
		Cog:       results,
		Creatures: killStatistics,
	})

	if presenterErr != nil {
		return presenterErr
	}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}

func (h *DetailHandler) getCogInformation(itemName string) (types.Cog, error) {
	var filters []types.Filter

	filters = append(filters, types.Filter{
		Name:    "name",
		Operand: constants.Equal,
		Value:   itemName,
	})

	criteria := types.Criteria{Filters: filters}

	result, err := h.cogRepository.FindOne(criteria)

	if err != nil {
		return types.Cog{}, err
	}

	return result, nil
}

func (h *DetailHandler) getKillStatistics(cog types.Cog) ([]types.CreatureKillStatistic, error) {
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

		result, err := h.killStatisticRepository.FindOne(criteria)

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
