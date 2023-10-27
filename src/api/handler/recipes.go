package handler

import (
	"fmt"
	"ncbs/api/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

var recipesTable = "Recipes"

type RecipeHandler struct {
	Database dynamodb.DynamoDB
}

func NewRecipeHandler(database dynamodb.DynamoDB) *RecipeHandler {
	return &RecipeHandler{
		Database: database,
	}
}

func (h *RecipeHandler) CreateRecipe(recipe model.Recipe) error {
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

	_, err := h.Database.PutItem(input)
	if err != nil {
		return fmt.Errorf("Could not put item into DynamoDB table: %w", err)
	}

	return nil
}

func (h *RecipeHandler) GetAllRecipes() ([]model.Recipe, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(recipesTable),
	}
	result, err := h.Database.Scan(input)
	if err != nil {
		return nil, fmt.Errorf("Could not scan DynamoDB table: %w", err)
	}

	var recipes []model.Recipe
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recipes)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal Query result items, %w", err)
	}

	return recipes, nil
}
