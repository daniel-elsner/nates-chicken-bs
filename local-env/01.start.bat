@echo off

set AWS_PROFILE=%1

if "%AWS_PROFILE%"=="" (
    set AWS_PROFILE=local-development
)

echo Using AWS Profile: '%AWS_PROFILE%' (Not important, just needs to exist) 

echo Starting Docker Compose...
docker-compose up -d

echo Waiting for DynamoDB to initialize...
timeout /t 10

echo Checking if Recipes table exists...
aws dynamodb list-tables --endpoint-url http://localhost:8000 | find "Recipes" >nul
if errorlevel 1 (
    echo Creating Recipes table...
    aws dynamodb create-table ^
    --table-name Recipes ^
    --attribute-definitions AttributeName=RecipeId,AttributeType=S ^
    --key-schema AttributeName=RecipeId,KeyType=HASH ^
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 ^
    --endpoint-url http://localhost:8000
) else (
    echo Recipes table already exists. Skipping creation.
)

echo Setup complete. 

echo Following docker-compose logs...
docker-compose logs -f