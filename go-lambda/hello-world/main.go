package main

import (
	"context"
	"hello-world/infra"
	"hello-world/repository"
	"hello-world/usecase"
	"log"

	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var dynamoClient *dynamodb.Client
var tableName string
var echoLambda *echoadapter.EchoLambda

type User struct {
	ID string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}


func main() {
	ctx := context.Background()

	dynamoClient, err := infra.NewDynamoClient(ctx)
	if err != nil {
		log.Fatalf("failed to create dynamo client: %v", err)
	}

	tableName = os.Getenv("TABLE_NAME")
	repo := repository.NewUserRepository(dynamoClient, tableName)
	uc := usecase.NewUserUsecase(repo)

	e := echo.New()
	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users/:id", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	echoLambda = echoadapter.New(e)
	lambda.Start(echoLambda.ProxyWithContext)
}
