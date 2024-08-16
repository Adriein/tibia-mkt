package service

import (
	"database/sql"
	"github.com/adriein/tibia-mkt/internal/tibia-mkt/repository"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"time"
)

type Seeder struct {
	csvRepository types.GoodRecordRepository
	cogRepository types.Repository[types.Good]
	cron          *DataSnapshotCron
	database      *sql.DB
}

func NewSeeder(
	csvRepository types.GoodRecordRepository,
	cogRepository types.Repository[types.Good],
	cron *DataSnapshotCron,
	database *sql.DB,
) *Seeder {
	return &Seeder{
		csvRepository: csvRepository,
		cogRepository: cogRepository,
		cron:          cron,
		database:      database,
	}
}

func (s *Seeder) Execute(request types.SeedRequest) error {
	for _, item := range request.Items {
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
	}

	for _, item := range request.Items {
		repo := repository.NewPgGoodRecordRepository(s.database, item.Name)

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

			if cronErr := s.cron.Execute(); cronErr != nil {
				return cronErr
			}
		}
	}

	return nil
}
