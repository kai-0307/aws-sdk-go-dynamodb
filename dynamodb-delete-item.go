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
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamodb.AttributeValue{
			"user_id": &dynamodb.AttributeValueMemberS{Value: "0"},
		},
	}

	_, err := svc.DeleteItem(context.TODO(), deleteItemInput)
	if err != nil {
		log.Fatalf("Failed to delete item: %v", err)
	}

	fmt.Println("Item deleted successfully!")
}
