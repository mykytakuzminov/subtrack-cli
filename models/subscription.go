package models

import (
	"fmt"
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

func (s Subscription) String() string {
	return fmt.Sprintf(
		"%-20s | %8.2f | %-10s | %s",
		s.Name,
		s.Price,
		s.BillingCycle,
		s.CreatedAt.Format("2006-01-02"))
}
