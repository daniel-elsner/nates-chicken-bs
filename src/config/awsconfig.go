package config

import (
	"os"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// GetAWSConfig returns an AWS configuration based on environment variables or defaults.
func GetAWSConfig() *aws.Config {
	var awsConfig *aws.Config

	// profile is expected to only be set when running locally, when deployed to AWS
	// we will be using the IAM role attached to AppRunner
	if profile := os.Getenv("AWS_PROFILE"); profile != "" {
		log.Printf("STARTUP LOG: Creating config with AWS profile %w", profile)
		awsConfig = &aws.Config{
			Region:      aws.String("us-east-2"),
			Credentials: credentials.NewSharedCredentials("", profile),
		}
	} else {
		awsConfig = &aws.Config{
			Region: aws.String("us-east-2"),
		}
	}

	return awsConfig
}
