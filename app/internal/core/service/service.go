package service

import "github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"

type ItemsService struct {
}

func NewItemsService() *ItemsService {
	return &ItemsService{}
}

func (svc *ItemsService) GetItems() ([]*domain.Item, error) {
	return nil, nil
}
func (svc *ItemsService) CreateItem(*domain.Item) (*domain.Item, error) {
	return nil, nil
}

// Throws:
// ErrItemNotFound if id not found
func (svc *ItemsService) GetItem(string) (*domain.Item, error) {
	return nil, nil
}

// Throws:
// ErrItemNotFound if id not found
func (svc *ItemsService) UpdateItem(string) (*domain.Item, error) {
	return nil, nil
}

// Throws:
// ErrItemNotFound if id not found
func (svc *ItemsService) DeleteItem(string) error {
	return nil
}
