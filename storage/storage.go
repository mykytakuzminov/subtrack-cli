package storage

import (
	"github.com/mykytakuzminov/subtrack-cli/models"
)

type Storage interface {
	Save(s models.Subscription) error
	GetAll() ([]models.Subscription, error)
	Delete(id string) error
}
