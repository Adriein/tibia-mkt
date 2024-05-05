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

func (s *Seeder) Execute() {

}
