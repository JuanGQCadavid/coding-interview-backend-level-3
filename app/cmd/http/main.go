package main

import (
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/ports"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/service"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/handlers"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/repositories/mysqlrepo"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/gin-gonic/gin"
)

const (
	userEnvName   = "userDB"
	passEnvName   = "passDB"
	dbnameEnvName = "dbnameDB"
	urlEnvName    = "urlDB"
)

var (
	router = gin.Default()
)

func init() {
	var (
		repo        ports.Repository
		svc         ports.Service
		httpHandler *handlers.HttpHandler
	)

	repo, err := mysqlrepo.NewRDSRepoFromEnv(userEnvName, passEnvName, dbnameEnvName, urlEnvName)
	if err != nil {
		logs.Error.Fatalln("Unable to start service, repository fail to be created", err.Error())
	}
	svc = service.NewItemsService(repo)

	httpHandler = handlers.NewHttpHandler(svc)
	httpHandler.SetRouter(router)
}

func main() {
	router.Run(":8000")
}
