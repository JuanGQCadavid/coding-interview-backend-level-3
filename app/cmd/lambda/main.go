package main

import (
	"context"

	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/ports"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/service"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/handlers"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/repositories/mysqlrepo"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

const (
	userEnvName   = "userDB"
	passEnvName   = "passDB"
	dbnameEnvName = "dbnameDB"
	urlEnvName    = "urlDB"
)

var (
	ginLambda *ginadapter.GinLambda
)

func init() {
	var (
		repo        ports.Repository
		svc         ports.Service
		httpHandler *handlers.HttpHandler
		router      = gin.Default()
	)

	repo, err := mysqlrepo.NewRDSRepoFromEnv(userEnvName, passEnvName, dbnameEnvName, urlEnvName)
	if err != nil {
		logs.Error.Fatalln("Unable to start service, repository fail to be created", err.Error())
	}
	svc = service.NewItemsService(repo)

	httpHandler = handlers.NewHttpHandler(svc)
	httpHandler.SetRouter(router)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
