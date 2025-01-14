package mysqlrepo

import (
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type SqlRepo struct {
	connector *PostgresSQLConnnector
}

func NewRDSRepoFromEnv() (*SqlRepo, error) {
	connector, err := NewConectorFromEnv()
	if err != nil {
		return nil, err
	}

	return &SqlRepo{
		connector: connector,
	}, nil
}

func (repo *SqlRepo) FetchOne(string) (*domain.Item, error) {
	return nil, nil
}
func (repo *SqlRepo) FetchAll() ([]*domain.Item, error) {
	return nil, nil
}
func (repo *SqlRepo) Create(item *domain.Item) error {
	logs.Info.Printf("Create Item %v\n", item)

	if results := repo.connector.session.Create(item); results.Error != nil {
		logs.Error.Println("An error ocoured!: ", results.Error)
		return results.Error
	}

	return nil
}
func (repo *SqlRepo) Update(*domain.Item) error {
	return nil
}
func (repo *SqlRepo) Delete(string) error {
	return nil
}
