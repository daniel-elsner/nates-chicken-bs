package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner"
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

	// Create ECR repository
	ecrRepo := awsecr.NewRepository(stack, jsii.String("ncbs-images"), &awsecr.RepositoryProps{
		RemovalPolicy:  awscdk.RemovalPolicy_DESTROY,
		RepositoryName: jsii.String("ncbs-images"),
	})

	// Create IAM Role
	role := awsiam.NewRole(stack, jsii.String("AppRunnerECRAccessRoleCDK"), &awsiam.RoleProps{
		// AssumedBy: awsiam.NewServicePrincipal(jsii.String("tasks.apprunner.amazonaws.com"), nil),
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("build.apprunner.amazonaws.com"), nil),
	})

	// Add policy to IAM Role for ECR access
	ecrRepo.GrantPull(role)

	awsapprunner.NewCfnService(stack, jsii.String("ncbs"), &awsapprunner.CfnServiceProps{
		ServiceName: jsii.String("ncbs"),
		SourceConfiguration: &awsapprunner.CfnService_SourceConfigurationProperty{
			AuthenticationConfiguration: &awsapprunner.CfnService_AuthenticationConfigurationProperty{
				AccessRoleArn: jsii.String(*role.RoleArn()),
			},
			AutoDeploymentsEnabled: jsii.Bool(true),
			ImageRepository: &awsapprunner.CfnService_ImageRepositoryProperty{
				ImageIdentifier:     jsii.String(*ecrRepo.RepositoryUri()),
				ImageRepositoryType: jsii.String("ECR"),
			},
		},
	})

	// // Create DynamoDB table
	// dynamoTable := awsdynamodb.NewTable(stack, jsii.String("MyTable"), &awsdynamodb.TableProps{
	// 	PartitionKey: &awsdynamodb.Attribute{
	// 		Name: jsii.String("id"),
	// 		Type: awsdynamodb.AttributeType_STRING,
	// 	},
	// })

	// // Grant permissions
	// appRunnerService.GrantPrincipal().AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
	// 	Actions:   jsii.Strings("dynamodb:*"),
	// 	Resources: jsii.Strings(dynamoTable.TableArn()),
	// }))

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
