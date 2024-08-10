package handler

import (
	"github.com/adriein/tibia-mkt/internal/cron"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type DataSnapshotHandler struct {
	cron *cron.DataSnapshotCron
}

func NewDataSnapshotHandler(
	cron *cron.DataSnapshotCron,
) *DataSnapshotHandler {
	return &DataSnapshotHandler{
		cron: cron,
	}
}

func (h *DataSnapshotHandler) Handler(w http.ResponseWriter, _ *http.Request) error {

	if cronErr := h.cron.Execute(); cronErr != nil {
		return cronErr
	}

	response := types.ServerResponse{Ok: true}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
