package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
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
	awsecr.NewRepository(stack, jsii.String("ncbs-images"), &awsecr.RepositoryProps{
		RemovalPolicy:  awscdk.RemovalPolicy_DESTROY,
		RepositoryName: jsii.String("ncbs-images"),
	})

	// Create App Runner service
	// appRunnerService := awsapprunner.NewService(stack, jsii.String("MyAppRunnerService"), &awsapprunner.ServiceProps{
	// 	// Configure source from ECR
	// 	Source: awsapprunner.Source.FromEcr(&awsapprunner.EcrProps{
	// 		Repository: ecrRepo,
	// 		// Additional options...
	// 	}),
	// 	// Additional options...
	// })

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

	log.Println("Using account:", account)
	log.Println("Using region:", region)

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
