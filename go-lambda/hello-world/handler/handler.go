package handler

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var dynamoClient *dynamodb.Client
var echoLambda *echoadapter.EchoLambda

var tableName = os.Getenv("TABLE_NAME")

type User struct {
	ID string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}

