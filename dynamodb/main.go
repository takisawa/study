package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "ClimbingGeos"

func main() {
	dynamoSession, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamodb.New(dynamoSession)

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String("CLIMBING1"),
			},
		},
		KeyConditionExpression: aws.String("ClimbingId = :id"),
	}

	res, err := db.Query(queryInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", res)
}
