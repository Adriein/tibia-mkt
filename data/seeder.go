package data

import "github.com/adriein/tibia-mkt/pkg/types"

type Seeder struct {
	repository types.CogRepository
}

func New(repository types.CogRepository) *Seeder {
	return &Seeder{
		repository: repository,
	}
}

func (s *Seeder) Execute() error {
	results, repositoryErr := s.repository.Find(types.Criteria{})

	if repositoryErr != nil {
		return repositoryErr
	}

	return nil
}
