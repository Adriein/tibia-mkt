package service

import (
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"time"
)

type SeederService struct {
	csvRepository       types.GoodRecordRepository
	cogRepository       types.Repository[types.Good]
	dataSnapshotService *DataSnapshotService
	factory             *helper.RepositoryFactory
}

func NewSeederService(
	csvRepository types.GoodRecordRepository,
	cogRepository types.Repository[types.Good],
	cron *DataSnapshotService,
	factory *helper.RepositoryFactory,
) *SeederService {
	return &SeederService{
		csvRepository:       csvRepository,
		cogRepository:       cogRepository,
		dataSnapshotService: cron,
		factory:             factory,
	}
}

func (s *SeederService) Execute(goodsToSeed []types.SeedGood) error {
	for _, item := range goodsToSeed {
		creatures := make([]types.GoodDrop, len(item.Creatures))

		for index, creature := range item.Creatures {
			creatures[index] = types.GoodDrop{Name: creature.Name, DropRate: creature.DropRate}
		}

		id := uuid.New()

		cog := types.Good{
			Id:        id.String(),
			Name:      item.Name,
			Link:      item.Wiki,
			Drop:      creatures,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		if saveErr := s.cogRepository.Save(cog); saveErr != nil {
			return saveErr
		}

		/*repo := s.factory.Get(item.Name)

		var filters []types.Filter

		filters = append(filters, types.Filter{Name: item.Name, Operand: "=", Value: item.Name})

		results, csvErr := s.csvRepository.Find(types.Criteria{Filters: filters})

		if csvErr != nil {
			return csvErr
		}

		for _, result := range results {
			if pgErr := repo.Save(result); pgErr != nil {
				return pgErr
			}

			if cronErr := s.dataSnapshotService.Execute(); cronErr != nil {
				return cronErr
			}
		}*/
	}

	return nil
}
