package handler

import (
	"encoding/json"
	service2 "github.com/adriein/tibia-mkt/internal/tibia-mkt/service"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type SeedHandler struct {
	service *service2.Seeder
}

func NewSeedHandler(
	service *service2.Seeder,
) *SeedHandler {
	return &SeedHandler{
		service: service,
	}
}

func (h *SeedHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	var request types.SeedRequest

	if decodeErr := json.NewDecoder(r.Body).Decode(&request); decodeErr != nil {
		return types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Handler",
			File:     "seed.go",
		}
	}

	if serviceErr := h.service.Execute(request); serviceErr != nil {
		return serviceErr
	}

	response := types.ServerResponse{
		Ok: true,
	}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
