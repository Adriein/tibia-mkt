package price

type Service struct {
	repository PriceRepository
}

func NewService(repository PriceRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetPrice(world string, good string) ([]*Price, error) {
	price, err := s.repository.FindByNameAndWorld(world, good)

	if err != nil {
		return nil, err
	}

	return price, nil
}
