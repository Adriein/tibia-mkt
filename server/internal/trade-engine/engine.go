package trade_engine

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type TradeEngine[T any] struct {
	config            *TradeEngineConfig
	repositoryFactory *service.RepositoryFactory
	algorithm         types.TradeEngineAlgorithm[T]
}

func NewTradeEngine[T any](
	factory *service.RepositoryFactory,
	config *TradeEngineConfig,
	algorithm types.TradeEngineAlgorithm[T],
) *TradeEngine[T] {
	return &TradeEngine[T]{
		config:            config,
		repositoryFactory: factory,
		algorithm:         algorithm,
	}
}

func (te *TradeEngine[T]) Execute(interval types.CogInterval) (types.TradeEngineResult, error) {
	cogs, retrieveCogsErr := te.retrieveCogInInterval(interval)

	if retrieveCogsErr != nil {
		return types.TradeEngineResult{}, retrieveCogsErr
	}

	_, err := te.algorithm.Apply(cogs)

	if err != nil {
		return types.TradeEngineResult{}, err
	}

	return types.TradeEngineResult{}, nil
}

func (te *TradeEngine[T]) retrieveCogInInterval(interval types.CogInterval) ([]types.CogSku, error) {
	var filters []types.Filter

	filters = append(
		filters,
		types.Filter{Name: "date", Operand: constants.GreaterThanOrEqual, Value: interval.From.Format(time.DateOnly)},
		types.Filter{Name: "date", Operand: constants.LessThanOrEqual, Value: interval.To.Format(time.DateOnly)},
	)

	criteria := types.Criteria{Filters: filters}

	cogRepository := te.repositoryFactory.Get(interval.Name)

	results, err := cogRepository.Find(criteria)

	if err != nil {
		return []types.CogSku{}, err
	}

	return results, nil
}
