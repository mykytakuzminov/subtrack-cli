package models

import (
	"testing"
)

func TestNewSubscription(t *testing.T) {
	const name string = "Spotify"
	const price float64 = 5.99
	const cycle string = "monthly"

	s := NewSubscription(name, price, cycle)

	if s.Name != name {
		t.Errorf("expected '%s', got '%s'", name, s.Name)
	}
	if s.Price != price {
		t.Errorf("expected '%.2f' got '%.2f'", price, s.Price)
	}
	if s.BillingCycle != cycle {
		t.Errorf("expected '%s' got '%s'", cycle, s.BillingCycle)
	}
}
