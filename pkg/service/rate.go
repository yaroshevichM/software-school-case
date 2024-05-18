package service

import (
	"errors"

	"github.com/yaroshevichM/software-school-case/pkg/repository"
)

type RateService struct {
	repo repository.Rate
}

func NewRateService(repo repository.Rate) *RateService {
	return &RateService{
		repo: repo,
	}
}

func (s *RateService) GetUSDtoUAHRate() (float64, error) {
	rates, err := s.repo.GetRates()
	if err != nil {
		return 0, err
	}

	for _, rate := range rates {
		if rate.Currency == "USD" && rate.BaseCurrency == "UAH" {
			return rate.Amount, nil
		}
	}

	return 0, errors.New("USD to UAH rate not found")
}
