package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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
	keyCond := expression.Key("user_id").Equal(expression.Value("3"))
	expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).Build()

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	queryOutput, err := svc.Query(context.TODO(), queryInput)
	if err != nil {
		log.Fatalf("Failed to query items: %v", err)
	}

	fmt.Printf("Query results: %v\n", queryOutput.Items)
}
