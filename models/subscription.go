package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID           string
	Name         string
	Price        float64
	BillingCycle string
	CreatedAt    time.Time
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
