package mysqlrepo

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresSQLConnnector struct {
	session *gorm.DB
}

func NewConector(dbUser string, dbPassword string, dbName string, dbUrl string) (*PostgresSQLConnnector, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?TimeZone=UTC", dbUser, dbPassword, dbUrl, dbName)
	session, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		logs.Error.Println("We fail to create the connection to the DB, error: ", err.Error())
	}
	return &PostgresSQLConnnector{
		session: session,
	}, nil

}

func NewConectorFromEnv(user, pass, dbname, url string) (*PostgresSQLConnnector, error) {
	dbUser, isPresentUser := os.LookupEnv(user)
	dbPassword, isPresentPass := os.LookupEnv(pass)
	dbName, isPresentName := os.LookupEnv(dbname)
	dbUrl, isPresentUrl := os.LookupEnv(url)

	if !isPresentUrl || !isPresentName || !isPresentPass || !isPresentUser {
		log.Println("dbUser: ", dbUser)
		log.Println("dbPassword: ", dbPassword)
		log.Println("dbName: ", dbName)
		log.Println("dbUrl: ", dbUrl)
		log.Fatalln("The ULR, Password or Username, dbName is not present in the env.")
	}

	return NewConector(dbUser, dbPassword, dbName, dbUrl)
}

func (conn *PostgresSQLConnnector) Migrate() {
	conn.session.AutoMigrate(&domain.Item{})
}
