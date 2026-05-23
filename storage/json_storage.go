package storage

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/mykytakuzminov/subtrack-cli/models"
)

type JSONStorage struct {
	filePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{filePath: filePath}
}

func (js *JSONStorage) Save(s models.Subscription) error {
	subscriptions, err := js.GetAll()
	if err != nil {
		return err
	}

	subscriptions = append(subscriptions, s)

	err = js.writeAll(subscriptions)
	if err != nil {
		return err
	}

	return nil
}

func (js *JSONStorage) GetAll() ([]models.Subscription, error) {
	data, err := os.ReadFile(js.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return []models.Subscription{}, nil
	}
	if err != nil {
		return nil, err
	}

	var subscriptions []models.Subscription
	err = json.Unmarshal(data, &subscriptions)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (js *JSONStorage) Delete(id string) error {
	subscriptions, err := js.GetAll()
	if err != nil {
		return err
	}

	var updatedSubscriptions []models.Subscription
	for _, s := range subscriptions {
		if !strings.HasPrefix(s.ID, id) {
			updatedSubscriptions = append(updatedSubscriptions, s)
		}
	}

	err = js.writeAll(updatedSubscriptions)
	if err != nil {
		return err
	}

	return nil
}

func (js *JSONStorage) writeAll(subs []models.Subscription) error {
	data, err := json.Marshal(subs)
	if err != nil {
		return err
	}

	if err = os.WriteFile(js.filePath, data, 0644); err != nil {
		return err
	}

	return nil
}
