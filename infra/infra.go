package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type InfraStackProps struct {
	awscdk.StackProps
	EnvName string
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)
	envName := props.EnvName

	// create ecr repository, apprunner will build from these images
	ecrRepo := awsecr.NewRepository(stack, jsii.String("ncbs-images"), &awsecr.RepositoryProps{
		RepositoryName: jsii.String("ncbs-images"),
		RemovalPolicy:  awscdk.RemovalPolicy_DESTROY,
	})

	// iam role for apprunner to access ecr
	buildRoleName := jsii.String(envName + "-apprunner-ncbs-build")
	ecrRole := awsiam.NewRole(stack, buildRoleName, &awsiam.RoleProps{
		RoleName:  buildRoleName,
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("build.apprunner.amazonaws.com"), nil),
	})
	ecrRepo.GrantPull(ecrRole)

	// iam role for apprunner to access while running
	instanceRoleName := jsii.String(envName + "-apprunner-ncbs-tasks")
	appRunnerInstanceRole := awsiam.NewRole(stack, instanceRoleName, &awsiam.RoleProps{
		RoleName:  instanceRoleName,
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("tasks.apprunner.amazonaws.com"), nil),
	})

	// defining the app runner service
	appRunnerName := jsii.String(envName + "-ncbs")
	awsapprunner.NewCfnService(stack, appRunnerName, &awsapprunner.CfnServiceProps{
		ServiceName: appRunnerName,
		InstanceConfiguration: &awsapprunner.CfnService_InstanceConfigurationProperty{
			Cpu:             jsii.String("0.25 vCPU"),
			Memory:          jsii.String("0.5 GB"),
			InstanceRoleArn: jsii.String(*appRunnerInstanceRole.RoleArn()),
		},
		SourceConfiguration: &awsapprunner.CfnService_SourceConfigurationProperty{
			AuthenticationConfiguration: &awsapprunner.CfnService_AuthenticationConfigurationProperty{
				AccessRoleArn: jsii.String(*ecrRole.RoleArn()),
			},
			AutoDeploymentsEnabled: jsii.Bool(true),
			ImageRepository: &awsapprunner.CfnService_ImageRepositoryProperty{
				ImageIdentifier:     jsii.String(*ecrRepo.RepositoryUri() + ":latest"),
				ImageRepositoryType: jsii.String("ECR"),
			},
		},
	})

	// creating dynamo db table
	recipesTableName := jsii.String("Recipes") // will handle environment namespacing later
	recipesTable := awsdynamodb.NewTable(stack, recipesTableName, &awsdynamodb.TableProps{
		TableName: recipesTableName,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("ID"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	recipesTable.GrantReadWriteData(appRunnerInstanceRole)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	envName := getRequiredEnvVar("CDK_ENV_NAME")

	// https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	NewInfraStack(app, "InfraStack", &InfraStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String(getRequiredEnvVar("CDK_DEPLOY_ACCOUNT")),
				Region:  jsii.String(getRequiredEnvVar("CDK_DEPLOY_REGION")),
			},
		},
		envName,
	})

	app.Synth(nil)
}

func getRequiredEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Required environment variable %s not set", name)
	}

	log.Println("Retrieved value ", value, " (from ", name, " env var)")

	return value
}
