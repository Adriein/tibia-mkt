package event

type Service struct {
	repository EventRepository
}

func NewService(repository EventRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetEvents(world string, good string) ([]*Event, error) {
	events, findErr := s.repository.FindByWorldAndGood(world, good)

	if findErr != nil {
		return nil, findErr
	}

	return events, nil
}
