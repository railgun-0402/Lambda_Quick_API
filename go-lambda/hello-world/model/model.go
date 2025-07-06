package main

type TableItem struct {
	ID string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}