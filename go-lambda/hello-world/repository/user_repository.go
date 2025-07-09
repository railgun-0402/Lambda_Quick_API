package repository

import (
	"context"
	"hello-world/domain"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userRepository struct {
	client *dynamodb.Client
	tableName string
}

func NewUserRepository(client *dynamodb.Client, tableName string) UserRepository {
    return &userRepository{client: client, tableName: tableName}
}

func (r *userRepository) CreateUser(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func (r *userRepository) GetUser(c echo.Context) error {
	id := c.Param("id")
	res, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil || res.Item == nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	var user domain.User
	if err := attributevalue.UnmarshalMap(res.Item, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (r *userRepository) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.ID = id

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (r *userRepository) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	_, err := r.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted")
}
