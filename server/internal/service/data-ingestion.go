package service

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"time"
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

func (dis *DataIngestionService) Execute(dto types.GoodRecordDto) error {
	id := uuid.New()

	date, dateParseErr := time.Parse(constants.IncomingTimeFormat, dto.Date)

	if dateParseErr != nil {
		return types.ApiError{
			Msg:      dateParseErr.Error(),
			Function: "Execute -> time.Parse()",
			File:     "service/data-ingestion.go",
			Values:   []string{constants.IncomingTimeFormat, dto.Date},
		}
	}

	repository := dis.factory.Get(dto.ItemName)

	good := types.GoodRecord{
		Id:        id.String(),
		ItemName:  dto.ItemName,
		Date:      date,
		BuyPrice:  dto.BuyOffer,
		SellPrice: dto.SellOffer,
		World:     dto.World,
	}

	if saveErr := repository.Save(good); saveErr != nil {
		return saveErr
	}

	return nil
}
