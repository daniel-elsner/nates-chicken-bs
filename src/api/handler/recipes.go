package handler

import (
	"log"
	"ncbs/api/model"
	"ncbs/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

var recipesTable = "Recipes"

func CreateRecipe(recipe model.Recipe) error {
	prepSteps := make([]*dynamodb.AttributeValue, len(recipe.PrepSteps))
	for i, step := range recipe.PrepSteps {
		prepSteps[i] = &dynamodb.AttributeValue{S: aws.String(step)}
	}

	cookSteps := make([]*dynamodb.AttributeValue, len(recipe.CookSteps))
	for i, step := range recipe.CookSteps {
		cookSteps[i] = &dynamodb.AttributeValue{S: aws.String(step)}
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ID":        {S: aws.String(uuid.New().String())},
			"name":      {S: aws.String(recipe.Name)},
			"prepSteps": {L: prepSteps},
			"cookSteps": {L: cookSteps},
		},
		TableName: aws.String(recipesTable),
	}

	_, err := db.DynamoDBClient.PutItem(input)
	if err != nil {
		log.Fatalf("Could not put item into DynamoDB table: %s", err)
		return err
	}

	return nil
}

func GetAllRecipes() ([]model.Recipe, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(recipesTable),
	}
	result, err := db.DynamoDBClient.Scan(input)
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
