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
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	scanOutput, err := svc.Scan(context.TODO(), scanInput)
	if err != nil {
		log.Fatalf("Failed to scan items: %v", err)
	}

	fmt.Printf("Scan results: %v\n", scanOutput.Items)
}
