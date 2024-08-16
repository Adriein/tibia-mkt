package handler

import (
	service2 "github.com/adriein/tibia-mkt/internal/tibia-mkt/service"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type DataSnapshotHandler struct {
	cron *service2.DataSnapshotCron
}

func NewDataSnapshotHandler(
	cron *service2.DataSnapshotCron,
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

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
