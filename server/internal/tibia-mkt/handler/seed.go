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
	Name      string      `json:"name"`
	Wiki      string      `json:"wiki"`
	Creatures []Creatures `json:"creatures"`
}

type Creatures struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"dropRate"`
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
		return types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Handler",
			File:     "seed.go",
		}
	}

	for _, item := range request.Items {
		repository := h.repositoryFactory.Get(item.Name)
		seeder := service.NewSeeder(h.csvRepository, repository)

		id := uuid.New()

		creatures := make([]types.CogCreature, len(item.Creatures))

		for index, creature := range item.Creatures {
			creatures[index] = types.CogCreature{Name: creature.Name, DropRate: creature.DropRate}
		}

		cog := types.Cog{
			Id:        id.String(),
			Name:      item.Name,
			Link:      item.Wiki,
			Creatures: creatures,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		if saveErr := h.cogRepository.Save(cog); saveErr != nil {
			return saveErr
		}

		if seederErr := seeder.Execute(item.Name); seederErr != nil {
			return seederErr
		}
	}

	response := types.ServerResponse{
		Ok: true,
	}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
