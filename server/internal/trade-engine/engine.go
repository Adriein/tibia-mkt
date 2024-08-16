package trade_engine

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type TradeEngine[T any] struct {
	config            *TradeEngineConfig
	repositoryFactory *helper.RepositoryFactory
	algorithm         types.TradeEngineAlgorithm[T]
}

func NewTradeEngine[T any](
	factory *helper.RepositoryFactory,
	config *TradeEngineConfig,
	algorithm types.TradeEngineAlgorithm[T],
) *TradeEngine[T] {
	return &TradeEngine[T]{
		config:            config,
		repositoryFactory: factory,
		algorithm:         algorithm,
	}
}

func (te *TradeEngine[T]) Execute(interval types.GoodRecordInterval) (T, error) {
	var failedResponse T
	cogs, retrieveCogsErr := te.retrieveCogInInterval(interval)

	if retrieveCogsErr != nil {
		return failedResponse, retrieveCogsErr
	}

	result, err := te.algorithm.Apply(cogs)

	if err != nil {
		return failedResponse, err
	}

	return result, nil
}

func (te *TradeEngine[T]) retrieveCogInInterval(interval types.GoodRecordInterval) ([]types.GoodRecord, error) {
	var filters []types.Filter

	filters = append(
		filters,
		types.Filter{Name: "date", Operand: constants.GreaterThanOrEqual, Value: interval.From},
		types.Filter{Name: "date", Operand: constants.LessThanOrEqual, Value: interval.To},
	)

	criteria := types.Criteria{Filters: filters}

	cogRepository := te.repositoryFactory.Get(interval.Name)

	results, err := cogRepository.Find(criteria)

	if err != nil {
		return []types.GoodRecord{}, err
	}

	return results, nil
}
