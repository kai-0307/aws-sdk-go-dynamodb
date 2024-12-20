package main

import (
	"context"
	"encoding/json"
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
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamodb.AttributeValue{
			"user_id": &dynamodb.AttributeValueMemberS{Value: "0"},
		},
	}

	getItemOutput, err := svc.GetItem(context.TODO(), getItemInput)
	if err != nil {
		log.Fatalf("Failed to get item: %v", err)
	}

	itemJSON, _ := json.Marshal(getItemOutput.Item)
	fmt.Printf("Item retrieved: %s\n", string(itemJSON))
}
