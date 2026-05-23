package storage

import (
	"testing"

	"github.com/mykytakuzminov/subtrack-cli/models"
)

const name string = "Spotify"
const price float64 = 5.99
const cycle string = "monthly"

func newTestJSONStorage(t *testing.T) *JSONStorage {
	return NewJSONStorage(t.TempDir() + "/test.json")
}

func TestSave(t *testing.T) {
	storage := newTestJSONStorage(t)
	s := models.NewSubscription(name, price, cycle)

	if err := storage.Save(s); err != nil {
		t.Errorf("Save() returned error: %v", err)
	}

	subs, err := storage.GetAll()
	if err != nil {
		t.Errorf("GetAll() returned error: %v", err)
	}

	if len(subs) != 1 {
		t.Fatalf("expected 1 subscription, got %d", len(subs))
	}

	if subs[0].Name != name {
		t.Errorf("expected '%s', got '%s'", name, subs[0].Name)
	}
	if subs[0].Price != price {
		t.Errorf("expected '%.2f' got '%.2f'", price, subs[0].Price)
	}
	if subs[0].BillingCycle != cycle {
		t.Errorf("expected '%s' got '%s'", cycle, subs[0].BillingCycle)
	}
}

func TestSaveMultiple(t *testing.T) {
	storage := newTestJSONStorage(t)
	s := models.NewSubscription(name, price, cycle)

	for i := 0; i < 2; i++ {
		if err := storage.Save(s); err != nil {
			t.Errorf("Save() returned error: %v", err)
		}
	}

	subs, err := storage.GetAll()
	if err != nil {
		t.Errorf("GetAll() returned error: %v", err)
	}

	if len(subs) != 2 {
		t.Fatalf("expected 2 subscriptions, got %d", len(subs))
	}
}

func TestGetAllEmpty(t *testing.T) {
	storage := newTestJSONStorage(t)
	subs, err := storage.GetAll()

	if err != nil {
		t.Errorf("GetAll() returned unexpected error: %v", err)
	}
	if len(subs) != 0 {
		t.Errorf("expected 0 subscriptions, got %d", len(subs))
	}
}

func TestDelete(t *testing.T) {
	storage := newTestJSONStorage(t)
	s := models.NewSubscription(name, price, cycle)

	if err := storage.Save(s); err != nil {
		t.Errorf("Save() returned error: %v", err)
	}

	subs, err := storage.GetAll()
	if err != nil {
		t.Errorf("GetAll() returned error: %v", err)
	}

	if err = storage.Delete(subs[0].ID); err != nil {
		t.Errorf("Delete() returned error: %v", err)
	}

	subs, err = storage.GetAll()
	if err != nil {
		t.Errorf("GetAll() returned error: %v", err)
	}

	if len(subs) != 0 {
		t.Errorf("expected 0 subscriptions, got %d", len(subs))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	storage := newTestJSONStorage(t)

	if err := storage.Delete("hd63j8"); err != nil {
		t.Errorf("Delete() returned unexpected error: %v", err)
	}
}
