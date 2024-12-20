package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var svc *dynamodb.Client
var tableName = "sampleTable"

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc = dynamodb.NewFromConfig(cfg)
}


func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	tableName := "sampleTable"

	// アイテムを追加
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]dynamodb.AttributeValue{
			"user_id":  &dynamodb.AttributeValueMemberS{Value: "0"},
			"username": &dynamodb.AttributeValueMemberS{Value: "Tatsuki"},
		},
	}

	_, err = svc.PutItem(context.TODO(), putItemInput)
	if err != nil {
		log.Fatalf("Failed to add item: %v", err)
	}

	fmt.Println("Item added successfully!")
}
