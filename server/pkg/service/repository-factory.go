package service

import "github.com/adriein/tibia-mkt/pkg/types"

type RepositoryFactory struct {
	repositories map[string]types.GoodRecordRepository
}

func NewRepositoryFactory(repositories []types.GoodRecordRepository) *RepositoryFactory {
	repoMap := make(map[string]types.GoodRecordRepository)

	for _, repository := range repositories {
		repoMap[repository.GoodName()] = repository
	}

	return &RepositoryFactory{
		repositories: repoMap,
	}
}

func (r *RepositoryFactory) Get(repository string) types.GoodRecordRepository {
	return r.repositories[repository]
}
