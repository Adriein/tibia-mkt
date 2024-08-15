package cron

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"time"
)

type DataSnapshotCron struct {
	cogRepository          types.Repository[types.Good]
	dataSnapshotRepository types.Repository[types.DataSnapshot]
	service                *service.DetailService
}

func NewDataSnapshotCron(
	cogRepository types.Repository[types.Good],
	dataSnapshotRepository types.Repository[types.DataSnapshot],
	service *service.DetailService,
) *DataSnapshotCron {
	return &DataSnapshotCron{
		cogRepository:          cogRepository,
		dataSnapshotRepository: dataSnapshotRepository,
		service:                service,
	}
}

func (dsc *DataSnapshotCron) Execute() error {
	cogs, cogRepoErr := dsc.cogRepository.Find(types.Criteria{Filters: make([]types.Filter, 0)})

	if cogRepoErr != nil {
		return cogRepoErr
	}

	for _, cog := range cogs {
		id := uuid.New()
		result, serviceErr := dsc.service.Execute(cog.Name)

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

		if snapshotRepoErr := dsc.dataSnapshotRepository.Save(snapshot); snapshotRepoErr != nil {
			return snapshotRepoErr
		}

	}
	return nil
}
