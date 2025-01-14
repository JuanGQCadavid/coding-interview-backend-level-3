package mysqlrepo

import "github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"

type MysqlRepo struct {
}

func (repo *MysqlRepo) FetchOne(string) (*domain.Item, error) {
	return nil, nil
}
func (repo *MysqlRepo) FetchAll(string) ([]*domain.Item, error) {
	return nil, nil
}
func (repo *MysqlRepo) Create(*domain.Item) error {
	return nil
}
func (repo *MysqlRepo) Update(*domain.Item) error {
	return nil
}
func (repo *MysqlRepo) Delete(string) error {
	return nil
}
