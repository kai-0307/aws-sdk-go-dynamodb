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
	filterExpr := expression.Name("username").Equal(expression.Value("Name 4"))
	projExpr := expression.NamesList(expression.Name("username"))
	expr, _ := expression.NewBuilder().WithFilter(filterExpr).WithProjection(projExpr).Build()

	scanInput := &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	scanOutput, err := svc.Scan(context.TODO(), scanInput)
	if err != nil {
		log.Fatalf("Failed to scan items with condition: %v", err)
	}

	fmt.Printf("Conditioned Scan results: %v\n", scanOutput.Items)
}
