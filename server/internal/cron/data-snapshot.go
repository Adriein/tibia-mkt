package cron

import (
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type DataSnapshotCron struct {
	cogRepository          types.Repository[types.Cog]
	dataSnapshotRepository types.Repository[types.DataSnapshot]
	service                *service.DetailService
}

func NewDataSnapshotCron(
	cogRepository types.Repository[types.Cog],
	dataSnapshotRepository types.Repository[types.DataSnapshot],
	service *service.DetailService,
) *DataSnapshotCron {
	return &DataSnapshotCron{
		cogRepository:          cogRepository,
		dataSnapshotRepository: dataSnapshotRepository,
		service:                service,
	}
}

func (dsc *DataSnapshotCron) Execute() ([]types.DataSnapshot, error) {
	cogs, cogRepoErr := dsc.cogRepository.Find(types.Criteria{Filters: make([]types.Filter, 0)})

	if cogRepoErr != nil {
		return nil, cogRepoErr
	}

	for _, cog := range cogs {
		result, serviceErr := dsc.service.Execute(cog.Name)

		if serviceErr != nil {
			return nil, serviceErr
		}

		snapshot := types.DataSnapshot{
			Id:           ",",
			Cog:          cog.Name,
			StdDeviation: result.StdDeviation,
			Mean:         result.SellPriceMean,
			ExecutedBy:   "tibia-mkt",
			CreatedAt:    time.Now().Format(time.DateTime),
			UpdatedAt:    time.Now().Format(time.DateTime),
		}

		if snapshotRepoErr := dsc.dataSnapshotRepository.Save(snapshot); snapshotRepoErr != nil {
			return nil, snapshotRepoErr
		}

	}
	return nil, nil
}
