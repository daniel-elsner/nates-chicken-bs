package db

import (
	"ncbs/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// InitDynamoClient initializes the DynamoDB client. This should be called
// before any DynamoDB operations are performed.
func InitDynamoClient(appConfig *config.Configuration, awsConfig *aws.Config) *dynamodb.DynamoDB {

	localConfig := *awsConfig
	if appConfig.DeployEnv == "local" {
		endpoint := "http://localhost:8000"
		localConfig.Endpoint = &endpoint
	}

	sess := session.Must(session.NewSession(&localConfig))
	return dynamodb.New(sess)
}
