package awsconfig

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// GetAWSConfig returns an AWS configuration based on environment variables or defaults.
func GetAWSConfig() *aws.Config {
	var awsConfig *aws.Config
	if profile := os.Getenv("AWS_PROFILE"); profile != "" {
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
