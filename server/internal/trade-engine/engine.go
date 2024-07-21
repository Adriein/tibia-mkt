package trade_engine

import (
	"fmt"
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

	result, err := te.algorithm.Apply(cogs)

	fmt.Println(result)

	if err != nil {
		return types.TradeEngineResult{}, err
	}

	return types.TradeEngineResult{}, nil
}

func (te *TradeEngine[T]) retrieveCogInInterval(interval types.CogInterval) ([]types.CogSku, error) {
	var filters []types.Filter

	filters = append(
		filters,
		types.Filter{Name: "date", Operand: ">=", Value: interval.From.Format(time.DateOnly)},
		types.Filter{Name: "date", Operand: "<=", Value: interval.To.Format(time.DateOnly)},
	)

	criteria := types.Criteria{Filters: filters}

	cogRepository := te.repositoryFactory.Get(interval.Name)

	results, err := cogRepository.Find(criteria)

	if err != nil {
		return []types.CogSku{}, err
	}

	return results, nil
}
