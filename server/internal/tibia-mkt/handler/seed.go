package handler

import (
	"encoding/json"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type SeedHandler struct {
	csvRepository types.CogRepository
	pgRepository  types.CogRepository
}

func NewSeedHandler(
	csvRepository types.CogRepository,
	pgRepository types.CogRepository,
) *SeedHandler {
	return &SeedHandler{
		csvRepository: csvRepository,
		pgRepository:  pgRepository,
	}
}

func (h *SeedHandler) Handler(w http.ResponseWriter, _ *http.Request) error {
	seeder := service.NewSeeder(h.csvRepository, h.pgRepository)

	if seederErr := seeder.Execute(); seederErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		response := &types.ServerResponse{
			Ok: false,
		}

		bytes, jsonErr := json.Marshal(response)

		if jsonErr != nil {
			return types.ApiError{
				Msg:      jsonErr.Error(),
				Function: "Handler",
				File:     "seed.go",
			}
		}

		_, writeErr := w.Write(bytes)

		if writeErr != nil {
			return types.ApiError{
				Msg:      writeErr.Error(),
				Function: "Handler",
				File:     "seed.go",
			}
		}

		return seederErr
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &types.ServerResponse{
		Ok: true,
	}

	bytes, jsonErr := json.Marshal(response)

	if jsonErr != nil {
		return types.ApiError{
			Msg:      jsonErr.Error(),
			Function: "Handler",
			File:     "seed.go",
		}
	}

	_, writeErr := w.Write(bytes)

	if writeErr != nil {
		return types.ApiError{
			Msg:      writeErr.Error(),
			Function: "Handler",
			File:     "seed.go",
		}
	}

	return nil
}
