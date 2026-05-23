package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	BillingCycle string    `json:"billing_cycle"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewSubscription(
	name string,
	price float64,
	billingCycle string,
) Subscription {
	return Subscription{
		ID:           uuid.NewString(),
		Name:         name,
		Price:        price,
		BillingCycle: billingCycle,
		CreatedAt:    time.Now(),
	}
}
