package handler

import (
	"encoding/json"
	"github.com/adriein/tibia-mkt/internal/service"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type DataIngestionHandler struct {
	service *service.DataIngestionService
}

func NewDataIngestionHandler(
	service *service.DataIngestionService,
) *DataIngestionHandler {
	return &DataIngestionHandler{
		service: service,
	}
}

func (h *DataIngestionHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	var request types.GoodRecordDto

	if decodeErr := json.NewDecoder(r.Body).Decode(&request); decodeErr != nil {
		return types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Handler",
			File:     "data-ingestion-handler.go",
		}
	}
	if serviceErr := h.service.Execute(request); serviceErr != nil {
		return serviceErr
	}

	response := types.ServerResponse{Ok: true}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
