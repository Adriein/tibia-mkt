package trade_engine

import (
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type TradeEngine struct {
	config            *TradeEngineConfig
	repositoryFactory service.RepositoryFactory
}

func NewTradeEngine(factory service.RepositoryFactory, config *TradeEngineConfig) *TradeEngine {
	return &TradeEngine{
		config:            config,
		repositoryFactory: factory,
	}
}

func (te *TradeEngine) Execute(interval types.CogInterval) (types.TradeEngineResult, error) {
	algorithm := te.config.Algorithm
	cogs, retrieveCogsErr := te.retrieveCogInInterval(interval)

	if retrieveCogsErr != nil {
		return types.TradeEngineResult{}, retrieveCogsErr
	}

	err := algorithm.Apply(cogs)

	if err != nil {
		return types.TradeEngineResult{}, err
	}

	return types.TradeEngineResult{}, nil
}

func (te *TradeEngine) retrieveCogInInterval(interval types.CogInterval) ([]types.CogSku, error) {
	var filters []types.Filter

	filters = append(
		filters,
		types.Filter{Name: "date", Operand: ">=", Value: interval.From},
		types.Filter{Name: "date", Operand: "<=", Value: interval.To},
	)

	criteria := types.Criteria{Filters: filters}

	cogRepository := te.repositoryFactory.Get(interval.Name)

	results, err := cogRepository.Find(criteria)

	if err != nil {
		return []types.CogSku{}, err
	}

	return results, nil
}
