package service

import (
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type DataIngestionService struct {
	factory *helper.RepositoryFactory
}

func NewDataIngestionService(
	factory *helper.RepositoryFactory,
) *DataIngestionService {
	return &DataIngestionService{
		factory: factory,
	}
}

func (dis *DataIngestionService) Execute(good types.GoodRecord) error {
	repository := dis.factory.Get(good.ItemName)

	if saveErr := repository.Save(good); saveErr != nil {
		return saveErr
	}

	return nil
}
