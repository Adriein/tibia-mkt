package handler

import (
	service2 "github.com/adriein/tibia-mkt/internal/tibia-mkt/service"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type KillStatisticsHandler struct {
	cron                    *service2.KillStatisticsCron
	cogRepository           types.Repository[types.Good]
	killStatisticRepository types.Repository[types.KillStatistic]
}

func NewKillStatisticsHandler(
	cron *service2.KillStatisticsCron,
	cogRepository types.Repository[types.Good],
	killStatisticRepository types.Repository[types.KillStatistic],
) *KillStatisticsHandler {
	return &KillStatisticsHandler{
		cron:                    cron,
		cogRepository:           cogRepository,
		killStatisticRepository: killStatisticRepository,
	}
}

func (h *KillStatisticsHandler) Handler(w http.ResponseWriter, _ *http.Request) error {
	cogs, cogErr := h.getTibiaMktTrackedCogs()

	if cogErr != nil {
		return cogErr
	}

	results, cronErr := h.cron.Execute(cogs)

	if cronErr != nil {
		return cronErr
	}

	for _, result := range results {
		if saveErr := h.killStatisticRepository.Save(result); saveErr != nil {
			return saveErr
		}
	}

	response := types.ServerResponse{Ok: true}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}

func (h *KillStatisticsHandler) getTibiaMktTrackedCogs() ([]types.Good, error) {
	var filters []types.Filter

	criteria := types.Criteria{Filters: filters}

	result, err := h.cogRepository.Find(criteria)

	if err != nil {
		return nil, err
	}

	return result, nil
}
