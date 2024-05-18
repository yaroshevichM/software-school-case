package models

type CreateSubscriptionInput struct {
	Email string `json:"email" db:"email"`
}

type Subscription struct {
	Id    int    `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}
