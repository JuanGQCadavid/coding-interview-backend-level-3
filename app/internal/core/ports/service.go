package ports

import (
	"errors"

	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"
)

var (
	ErrItemNotFound  = errors.New("err the item does not exist")
	ErrNegativePrice = errors.New("err the price should not be negative")
	ErrMissingPrice  = errors.New("err the price should exist")
	ErrInternalDB    = errors.New("err we found an error on DB")
)

type Service interface {
	GetItems() ([]*domain.Item, error)
	CreateItem(*domain.Item) (*domain.Item, error)

	// Throws:
	// ErrItemNotFound if id not found
	GetItem(string) (*domain.Item, error)

	// Throws:
	// ErrItemNotFound if id not found
	UpdateItem(*domain.Item) (*domain.Item, error)

	// Throws:
	// ErrItemNotFound if id not found
	DeleteItem(string) error
}
