package ports

import "github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"

type Repository interface {
	FetchOne(string) (*domain.Item, error)
	FetchAll() ([]*domain.Item, error)
	Create(*domain.Item) error
	Update(*domain.Item) error
	Delete(string) error
}
