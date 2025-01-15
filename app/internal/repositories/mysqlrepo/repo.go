package mysqlrepo

import (
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/gorm"
)

type SqlRepo struct {
	connector *PostgresSQLConnnector
}

func NewRDSWithConnector(connector *PostgresSQLConnnector) *SqlRepo {
	connector.Migrate()
	return &SqlRepo{
		connector: connector,
	}
}

func NewRDSRepoFromEnv(userEnvName, passEnvName, dbnameEnvName, urlEnvName string) (*SqlRepo, error) {
	connector, err := NewConectorFromEnv(userEnvName, passEnvName, dbnameEnvName, urlEnvName)
	if err != nil {
		return nil, err
	}
	connector.Migrate()
	return &SqlRepo{
		connector: connector,
	}, nil
}

func (repo *SqlRepo) FetchOne(id string) (*domain.Item, error) {

	var (
		item *domain.Item = &domain.Item{}
	)
	result := repo.connector.session.First(item, &id)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			logs.Warning.Println("Item not found in db ", id)
			return nil, nil
		}

		logs.Error.Println("Error while Fetching date: ", result.Error)
		return nil, result.Error
	}

	return item, nil
}
func (repo *SqlRepo) FetchAll() ([]*domain.Item, error) {
	var (
		items []*domain.Item
	)

	result := repo.connector.session.Find(&items)
	if result.Error != nil {

		logs.Error.Println("Error while Fetching date: ", result.Error)
		return nil, result.Error
	}
	return items, nil
}
func (repo *SqlRepo) Create(item *domain.Item) error {
	logs.Info.Printf("Create Item %v\n", item)

	if results := repo.connector.session.Create(item); results.Error != nil {
		logs.Error.Println("An error ocoured!: ", results.Error)
		return results.Error
	}

	return nil
}
func (repo *SqlRepo) Update(item *domain.Item) error {
	if results := repo.connector.session.Save(item); results.Error != nil {
		logs.Error.Println("An error ocoured while updating!: ", results.Error)
		return results.Error
	}
	return nil
}
func (repo *SqlRepo) Delete(id string) error {
	logs.Info.Println("Deleting ", id)
	if results := repo.connector.session.Unscoped().Delete(&domain.Item{}, &id); results.Error != nil {
		logs.Error.Println("An error ocoured!: ", results.Error)
		return results.Error
	}
	return nil
}
