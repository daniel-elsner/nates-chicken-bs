package handler

import (
	"log"
	"ncbs/api/awsconfig"
	"ncbs/api/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func InitDB() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(awsconfig.GetAWSConfig()))
	return dynamodb.New(sess)
}

func GetAllRecipes() ([]model.Recipe, error) {
	svc := InitDB()
	input := &dynamodb.ScanInput{
		TableName: aws.String("Recipes"), // your table name
	}
	result, err := svc.Scan(input)
	if err != nil {
		log.Fatalf("Could not scan DynamoDB table: %s", err)
		return nil, err
	}

	var recipes []model.Recipe
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recipes)
	if err != nil {
		log.Fatalf("Failed to unmarshal Query result items, %v", err)
		return nil, err
	}

	return recipes, nil
}
