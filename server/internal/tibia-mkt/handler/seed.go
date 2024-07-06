package handler

import (
	"encoding/json"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type SeedGood struct {
	Name string `json:"name"`
	Wiki string `json:"wiki"`
}

type SeedRequest struct {
	Items []SeedGood `json:"items"`
}

type SeedHandler struct {
	csvRepository     types.CogRepository
	repositoryFactory *service.RepositoryFactory
	cogRepository     types.Repository
}

func NewSeedHandler(
	csvRepository types.CogRepository,
	repositoryFactory *service.RepositoryFactory,
	cogRepository types.Repository,
) *SeedHandler {
	return &SeedHandler{
		csvRepository:     csvRepository,
		repositoryFactory: repositoryFactory,
		cogRepository:     cogRepository,
	}
}

func (h *SeedHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	var request SeedRequest

	if decodeErr := json.NewDecoder(r.Body).Decode(&request); decodeErr != nil {
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

		return decodeErr
	}

	for _, item := range request.Items {
		repository := h.repositoryFactory.Get(item.Name)
		seeder := service.NewSeeder(h.csvRepository, repository)

		id := uuid.New()

		cog := types.Cog{
			Id:        id.String(),
			Name:      item.Name,
			Link:      item.Wiki,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if saveErr := h.cogRepository.Save(cog); saveErr != nil {
			return types.ApiError{
				Msg:      saveErr.Error(),
				Function: "Handler",
				File:     "seed.go",
			}
		}

		if seederErr := seeder.Execute(item.Name); seederErr != nil {
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
