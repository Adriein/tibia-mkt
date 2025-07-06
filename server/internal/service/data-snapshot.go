package service

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"time"
)

type DataSnapshotService struct {
	cogRepository          types.Repository[types.Good]
	dataSnapshotRepository types.Repository[types.DataSnapshot]
	service                *DetailService
}

func NewDataSnapshotService(
	cogRepository types.Repository[types.Good],
	dataSnapshotRepository types.Repository[types.DataSnapshot],
	service *DetailService,
) *DataSnapshotService {
	return &DataSnapshotService{
		cogRepository:          cogRepository,
		dataSnapshotRepository: dataSnapshotRepository,
		service:                service,
	}
}

func (dss *DataSnapshotService) Execute() error {
	cogs, cogRepoErr := dss.cogRepository.Find(types.Criteria{Filters: make([]types.Filter, 0)})

	if cogRepoErr != nil {
		return cogRepoErr
	}

	for _, cog := range cogs {
		id := uuid.New()
		result, serviceErr := dss.service.Execute(cog.Name)

		totalDropped := 0

		if serviceErr != nil {
			return serviceErr
		}

		for _, creature := range result.Creatures {
			itemsDropped := float64(creature.KillStatistic) * (creature.DropRate / 100)

			totalDropped += int(itemsDropped)
		}

		snapshot := types.DataSnapshot{
			Id:           id.String(),
			Cog:          cog.Name,
			StdDeviation: result.StdDeviation,
			Mean:         result.SellPriceMean,
			TotalDropped: totalDropped,
			ExecutedBy:   constants.TibiaMktCronUser,
			CreatedAt:    time.Now().Format(time.DateTime),
			UpdatedAt:    time.Now().Format(time.DateTime),
		}

		if snapshotRepoErr := dss.dataSnapshotRepository.Save(snapshot); snapshotRepoErr != nil {
			return snapshotRepoErr
		}

	}
	return nil
}
