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
	updateExpr := expression.Set(expression.Name("username"), expression.Value("Akira"))
	expr, _ := expression.NewBuilder().WithUpdate(updateExpr).Build()

	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamodb.AttributeValue{
			"user_id": &dynamodb.AttributeValueMemberS{Value: "0"},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err := svc.UpdateItem(context.TODO(), updateItemInput)
	if err != nil {
		log.Fatalf("Failed to update item: %v", err)
	}

	fmt.Println("Item updated successfully!")
}
