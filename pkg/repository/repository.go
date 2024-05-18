package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yaroshevichM/software-school-case/pkg/models"
)

type Subscription interface {
	Create(createSubscription models.CreateSubscriptionInput) (int, error)
	GetAll() ([]models.Subscription, error)
	GetByEmail(email string) (models.Subscription, error)
}

type Rate interface {
	GetRates() ([]models.Rate, error)
}

type Repository struct {
	Subscription
	Rate
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Subscription: newSubscriptionRepository(db),
		Rate:         newRateRepository(db),
	}
}
