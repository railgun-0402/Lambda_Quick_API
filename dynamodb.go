package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type TableItem struct {
	ID string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}

type TableKey struct {
	ID string `dynamodbav:"id"`
}

func PutItem() {
	// AWS接続設定
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	// データの追加
	log.Println("put item")
	item := TableItem{
		ID: "user1",
		Name: "taro",
	}
	// DynamoDB用の形式（AttributeValueのマップ） に変換
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("golang-work-sample"),
		Item: av,
	})
	if err != nil {
		log.Fatal(err)
	}
}
