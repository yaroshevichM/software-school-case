package service

import (
	"github.com/yaroshevichM/software-school-case/pkg/models"
	"github.com/yaroshevichM/software-school-case/pkg/repository"
)

type Subscription interface {
	Create(createSubscription models.CreateSubscriptionInput) (int, error)
	GetAll() ([]models.Subscription, error)
	GetByEmail(email string) (models.Subscription, error)
}

type Rate interface {
	GetUSDtoUAHRate() (float64, error)
}

type Mail interface {
	SendMail(to, subject, body string) error
}

type Service struct {
	Subscription
	Rate
	Mail
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repo.Subscription),
		Rate:         NewRateService(repo.Rate),
		Mail:         NewMailService("smtp.example.com", 587, "your_username", "your_password"),
	}
}
