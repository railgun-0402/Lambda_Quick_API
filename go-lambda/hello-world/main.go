package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func init() {
	ctx := context.Background()
	var err error

	dynamoClient, err = NewDynamoClient(ctx)
	if err != nil {
		log.Fatalf("failed to create dynamo client: %v", err)
	}
	tableName = os.Getenv("TABLE_NAME")

	e := echo.New()
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	echoLambda = echoadapter.New(e)
}

func main() {
	lambda.Start(echoLambda.ProxyWithContext)
}

func createUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	res, err := dynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil || res.Item == nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	var user User
	if err := attributevalue.UnmarshalMap(res.Item, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.ID = id

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	_, err := dynamoClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted")
}
