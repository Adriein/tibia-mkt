package data

import "github.com/adriein/tibia-mkt/pkg/types"

type Seeder struct {
	csvRepository types.CogRepository
	pgRepository  types.CogRepository
}

func New(csvRepository types.CogRepository, pgRepository types.CogRepository) *Seeder {
	return &Seeder{
		csvRepository: csvRepository,
		pgRepository:  pgRepository,
	}
}

func (s *Seeder) Execute() error {
	results, csvErr := s.csvRepository.Find(types.Criteria{})

	if csvErr != nil {
		return csvErr
	}

	for _, result := range results {
		if pgErr := s.pgRepository.Save(result); pgErr != nil {
			return pgErr
		}
	}

	return nil
}
