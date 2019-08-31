package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

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

	res, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Persons"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String("1"),
			},
		},
		AttributesToGet: []*string{
			aws.String("Name"),
		},
		ConsistentRead:         aws.Bool(true),
		ReturnConsumedCapacity: aws.String("NONE"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", res)
}
