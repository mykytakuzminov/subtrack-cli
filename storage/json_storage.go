package storage

import (
	"github.com/mykytakuzminov/subtrack-cli/models"
)

type JSONStorage struct {
	filePath string
}

func NewJSONStorage(filePath string) JSONStorage {
	return JSONStorage{filePath: filePath}
}

func (js JSONStorage) Save(s models.Subscription) error {
	return nil
}

func (js JSONStorage) GetAll() ([]models.Subscription, error) {
	return make([]models.Subscription, 0), nil
}

func (js JSONStorage) Delete(id string) error {
	return nil
}
