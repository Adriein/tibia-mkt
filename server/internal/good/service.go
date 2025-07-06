package good

type Service struct {
	repository GoodRepository
}

func NewService(repository GoodRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetGood(itemName string) (*Good, error) {
	good, err := s.repository.FindByName(itemName)

	if err != nil {
		return nil, err
	}

	return good, nil
}
