package cron

import (
	"github.com/adriein/tibia-mkt/pkg/types"
)

type DataSnapshotCron struct{}

func NewDataSnapshotCron() *DataSnapshotCron {
	return &DataSnapshotCron{}
}

func (dsc *DataSnapshotCron) Execute() ([]types.DataSnapshot, error) {
	return nil, nil
}
