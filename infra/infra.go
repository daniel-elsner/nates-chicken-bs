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
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create ecr repository, apprunner will build from these images
	ecrRepo := awsecr.NewRepository(stack, jsii.String("ncbs-images"), &awsecr.RepositoryProps{
		RepositoryName: jsii.String("ncbs-images"),
		RemovalPolicy:  awscdk.RemovalPolicy_DESTROY,
	})

	// iam role for apprunner to access ecr
	buildRoleName := jsii.String("dev-apprunner-ncbs-build")
	ecrRole := awsiam.NewRole(stack, buildRoleName, &awsiam.RoleProps{
		RoleName:  buildRoleName,
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("build.apprunner.amazonaws.com"), nil),
	})
	ecrRepo.GrantPull(ecrRole)

	// iam role for apprunner to access while running
	instanceRoleName := jsii.String("dev-apprunner-ncbs-tasks")
	appRunnerInstanceRole := awsiam.NewRole(stack, instanceRoleName, &awsiam.RoleProps{
		RoleName:  instanceRoleName,
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("tasks.apprunner.amazonaws.com"), nil),
	})

	// defining the app runner service
	appRunnerName := jsii.String("dev-ncbs")
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

	NewInfraStack(app, "InfraStack", &InfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	account := os.Getenv("CDK_DEPLOY_ACCOUNT")
	region := os.Getenv("CDK_DEPLOY_REGION")

	log.Println("Using account:", account, " (from CDK_DEPLOY_ACCOUNT env var)")
	log.Println("Using region:", region, " (from CDK_DEPLOY_REGION env var)")

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
