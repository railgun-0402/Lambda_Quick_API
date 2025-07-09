package infra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// TODOこちらのコードを使用する
type TableItem struct {
	ID string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}

type ItemKey struct {
	ID string `dynamodbav:"id"`
}

func NewDynamoClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

// func client() *dynamodb.Client {
// 	// AWS接続設定
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return dynamodb.NewFromConfig(cfg)
// }

// func PutItem() {
// 	// AWS接続設定
// 	client := client()

// 	// データの追加
// 	log.Println("put item")
// 	item := TableItem{
// 		ID: "user1",
// 		Name: "taro",
// 	}
// 	// DynamoDB用の形式（AttributeValueのマップ） に変換
// 	av, err := attributevalue.MarshalMap(item)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
// 		TableName: aws.String("golang-work-sample"),
// 		Item: av,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func GetItem() {
// 	// AWS接続設定
// 	client := client()

// 	// データの追加
// 	log.Println("get item")
// 	itemKey := ItemKey{
// 		ID: "user1",
// 	}
// 	av, err := attributevalue.MarshalMap(itemKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
// 		TableName: aws.String("golang-work-sample"),
// 		Key: av,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 取得したItemを構造体に変換
// 	item := TableItem{}
// 	err = attributevalue.UnmarshalMap(res.Item, &item)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("id:", item.ID, "name:", item.Name)
// }

// func Update() {
// 	client := client()

// 	// データの更新
// 	log.Println("update item")
// 	item := TableItem{
// 		ID: "user1",
// 		Name: "jiro",
// 	}
// 	av, err := attributevalue.MarshalMap(item)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
// 		TableName: aws.String("golang-work-sample"),
// 		Item: av,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 更新データの取得
// 	log.Println("get updated item")
// 	itemKey := ItemKey{
// 		ID: "user1",
// 	}
// 	av, err = attributevalue.MarshalMap(itemKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
// 		TableName: aws.String("golang-work-sample"),
// 		Key: av,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	item = TableItem{}
// 	err = attributevalue.UnmarshalMap(res.Item, &item)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("id:", item.ID, "name:", item.Name)
// }

// func Delete () {
// 	client := client()

// 	// データの削除
// 	log.Println("delete item")
// 	itemKey := ItemKey{
// 		ID: "user1",
// 	}

// 	av, err := attributevalue.MarshalMap(itemKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
// 		TableName: aws.String("golang-work-sample"),
// 		Key: av,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
