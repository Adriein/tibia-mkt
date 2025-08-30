package price

import "github.com/adriein/tibia-mkt/pkg/constants"

type Service struct {
	repository PriceRepository
}

func NewService(repository PriceRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetPrices(world string, good string) ([]*Price, error) {
	var prices []*Price

	sellOfferResults, sellOfferErr := s.repository.FindOffersByGoodAndWorld(world, good, constants.SellOffer)

	if sellOfferErr != nil {
		return nil, sellOfferErr
	}

	prices = append(prices, sellOfferResults...)

	buyOfferResults, buyOfferErr := s.repository.FindOffersByGoodAndWorld(world, good, constants.BuyOffer)

	if buyOfferErr != nil {
		return nil, sellOfferErr
	}

	prices = append(prices, buyOfferResults...)

	return prices, nil
}
