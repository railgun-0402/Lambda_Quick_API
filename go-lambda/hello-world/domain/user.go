package domain

type User struct {
	ID string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}