package service

import (
	"github.com/adriein/tibia-mkt/pkg/types"
)

type SearchGoodService struct {
	goodRepository types.Repository[types.Good]
}

func NewSearchGoodService(
	goodRepository types.Repository[types.Good],
) *SearchGoodService {
	return &SearchGoodService{
		goodRepository: goodRepository,
	}
}

func (sgs *SearchGoodService) Execute() ([]types.Good, error) {
	goods, goodsRepoErr := sgs.goodRepository.Find(types.Criteria{Filters: make([]types.Filter, 0)})

	if goodsRepoErr != nil {
		return nil, goodsRepoErr
	}

	return goods, nil
}
