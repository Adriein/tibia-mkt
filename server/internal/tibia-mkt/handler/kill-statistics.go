package handler

import (
	"github.com/adriein/tibia-mkt/internal/cron"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type KillStatisticsHandler struct {
	cron          *cron.KillStatisticsCron
	cogRepository types.Repository
}

func NewKillStatisticsHandler(
	cron *cron.KillStatisticsCron,
	cogRepository types.Repository,
) *KillStatisticsHandler {
	return &KillStatisticsHandler{
		cron:          cron,
		cogRepository: cogRepository,
	}
}

func (h *KillStatisticsHandler) Handler(w http.ResponseWriter, _ *http.Request) error {
	cogs, cogErr := h.getTibiaMktTrackedCogs()

	if cogErr != nil {
		return cogErr
	}

	if cronErr := h.cron.Execute(cogs); cronErr != nil {
		return cronErr
	}

	response := types.ServerResponse{Ok: true}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}

func (h *KillStatisticsHandler) getTibiaMktTrackedCogs() ([]types.Cog, error) {
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
