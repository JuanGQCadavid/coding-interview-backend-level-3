package service

import (
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/google/uuid"
)

type ItemsService struct {
	repository ports.Repository
}

func NewItemsService(repository ports.Repository) *ItemsService {
	return &ItemsService{
		repository: repository,
	}
}

// Throws:
// ErrItemNotFound if id not found
func (svc *ItemsService) GetItem(id string) (*domain.Item, error) {
	item, err := svc.repository.FetchOne(id)

	if err != nil {
		logs.Error.Println("There is an error with the DB", err.Error())
		return nil, ports.ErrInternalDB
	}

	if item == nil || item.Id == "" {
		return nil, ports.ErrItemNotFound
	}

	return item, nil
}

func (svc *ItemsService) GetItems() ([]*domain.Item, error) {
	return svc.repository.FetchAll()
}

func (svc *ItemsService) genUUID() string {
	return uuid.NewString()
}

func (svc *ItemsService) validateRequest(item *domain.Item) error {

	if item.Price < 0 {
		return ports.ErrNegativePrice
	}

	return nil
}

func (svc *ItemsService) CreateItem(item *domain.Item) (*domain.Item, error) {
	if err := svc.validateRequest(item); err != nil {
		return nil, err
	}

	item.Id = svc.genUUID()

	if err := svc.repository.Create(item); err != nil {
		logs.Error.Println("There is an error on db: ", err.Error())
		return nil, ports.ErrInternalDB
	}

	return item, nil
}

// Throws:
// ErrItemNotFound if id not found
func (svc *ItemsService) UpdateItem(item *domain.Item) (*domain.Item, error) {
	if err := svc.validateRequest(item); err != nil {
		return nil, err
	}

	originalItem, _ := svc.repository.FetchOne(item.Id)

	if originalItem == nil {
		return nil, ports.ErrItemNotFound
	}

	if _, err := svc.UpdateItem(item); err != nil {
		logs.Error.Println("There is an error on db: ", err.Error())
		return nil, ports.ErrInternalDB
	}

	return item, nil
}

// Throws:
// ErrItemNotFound if id not found
func (svc *ItemsService) DeleteItem(id string) error {
	originalItem, _ := svc.repository.FetchOne(id)

	if originalItem == nil {
		return ports.ErrItemNotFound
	}

	if err := svc.DeleteItem(id); err != nil {
		logs.Error.Println("There is an error on db: ", err.Error())
		return ports.ErrInternalDB
	}

	return nil
}
