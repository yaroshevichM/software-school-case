package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/yaroshevichM/software-school-case/pkg/models"
)

type SubscriptionRepository struct {
	db *sqlx.DB
}

func newSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(item models.CreateSubscriptionInput) (int, error) {

	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createItemQuery := fmt.Sprintf(`INSERT INTO "%s" (email) values ($1) RETURNING id`, subscriptionTable)
	logrus.Print(createItemQuery)
	row := tx.QueryRow(createItemQuery, item.Email)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *SubscriptionRepository) GetAll() ([]models.Subscription, error) {
	var items []models.Subscription

	getAllItemsQuery := fmt.Sprintf(`SELECT * FROM "%s"`, subscriptionTable)
	err := r.db.Select(&items, getAllItemsQuery)
	logrus.Print(items)
	if err != nil {
		return items, err
	}

	return items, err
}

func (r *SubscriptionRepository) GetByEmail(email string) (models.Subscription, error) {
	var item models.Subscription

	getByEmailQuery := fmt.Sprintf(`SELECT * FROM "%s" WHERE email = $1`, subscriptionTable)
	err := r.db.Get(&item, getByEmailQuery, email)

	if err != nil {
		return item, err
	}

	return item, err
}
