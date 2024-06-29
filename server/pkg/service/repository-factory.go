package service

import "github.com/adriein/tibia-mkt/pkg/types"

type RepositoryFactory struct {
	repositories map[string]types.CogRepository
}

func NewRepositoryFactory(repositories []types.CogRepository) *RepositoryFactory {
	repoMap := make(map[string]types.CogRepository)

	for _, repository := range repositories {
		repoMap[repository.EntityName()] = repository
	}

	return &RepositoryFactory{
		repositories: repoMap,
	}
}

func (r *RepositoryFactory) Get(repository string) types.CogRepository {
	return r.repositories[repository]
}
