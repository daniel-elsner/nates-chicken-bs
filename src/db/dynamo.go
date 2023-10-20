package db

import (
	"ncbs/config"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoDBClient *dynamodb.DynamoDB

// InitDynamoClient initializes the DynamoDB client. This should be called
// before any DynamoDB operations are performed.
func InitDynamoClient() {
	sess := session.Must(session.NewSession(config.GetAWSConfig()))
	DynamoDBClient = dynamodb.New(sess)
}
