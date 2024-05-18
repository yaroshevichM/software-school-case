package service

import (
	"fmt"

	"github.com/yaroshevichM/software-school-case/pkg/models"
	"github.com/yaroshevichM/software-school-case/pkg/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(createSubscription models.CreateSubscriptionInput) (int, error) {
	var zeroSubscription models.Subscription
	existingEmail, err := s.repo.GetByEmail(createSubscription.Email)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return 0, err
	}

	if existingEmail != zeroSubscription {
		return 0, fmt.Errorf("email already exists")
	}

	return s.repo.Create(createSubscription)
}

func (s *SubscriptionService) GetAll() ([]models.Subscription, error) {
	return s.repo.GetAll()
}

func (s *SubscriptionService) GetByEmail(email string) (models.Subscription, error) {
	return s.repo.GetByEmail(email)
}
