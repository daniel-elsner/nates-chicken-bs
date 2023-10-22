# Notes Regarding Initial Infrastructure Setup

This is a dumping ground, not going to be organized. Just me trying to get the very initial setup done & the interactions between the code & AWS.

## AppRunner

Created a simple AppRunner instance which deploys automatically on new pushes to an ECR repo (handled by the GitHub Actions workflow). It's looking for Docker images with the `:latest` tag.

## DynamoDB

Initially created a `Recipes` table with `ID` as the partition key. Intentionally going as basic as possible to start, the real table would need to support querying by various attributes, which may not even make DynamoDB the best choice. 

## Doing Things Locally

Install the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html).

From that point, using named profiles seems like a nice way to go. 

First I created a group in AWS called `local-developers`, then created a new profile called
`local-daniel-elsner` and added it to the group. Initially the group has no permissions at all.

Then I created an access key for the user and downloaded it. I know doing access keys is not the ideal setup. I'm just not in the mood for figuring out all the advanced IAM setup, we can deal with this later.

Then I ran `aws configure --profile local-development` and entered the access key ID and secret access key. I left the default region as `us-east-2` and the default output format blank.

Then I ran `aws dynamodb scan --table-name Recipes --profile local-development` and get the following error (expected):

```
An error occurred (AccessDeniedException) when calling the Scan operation: User: arn:aws:iam::375842827666:user/local-daniel-elsner is not authorized to perform: dynamodb:Scan on resource: arn:aws:dynamodb:us-east-2:375842827666:table/Recipes because no identity-based policy allows the dynamodb:Scan action
```

Obviously I need to grant permissions to the user group. I manually did this by going to the `local-developers` group, Permissions tab, and clicking Add permissions -> Attach policies. From here I added the `AmazonDynamoDBFullAccess` permission.

Now when I query I get a result. I can then re-use this profile in the app running locally via something like this:

```go
func InitDB() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("", "local-development"),
	}))
	return dynamodb.New(sess)
}
```

This obviously sucks but it's a start. Note, if the local credentials file has a [default] profile set then we wouldn't even need to specify "local-development" here. Ultimately what I'm going to settle on is an environment variable `AWS_PROFILE` that specifies the profile name, and then the app will use that to load the credentials. 

```go
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
```

This way when the app is deployed to AppRunner it should just work through the IAM role assigned to the AppRunner instance, and not care at all about a local profile.

Note, I did test setting a local environment variable `AWS_PROFILE` and it worked but was oddly annoying. For Windows the only thing that wound up working was (via PowerShell):

```
$env:AWS_PROFILE="local-development"
```

Verify it by running:

```
Write-Host $env:AWS_PROFILE
```

This does begin to balloon the docker run command though:

```
docker run -v C:\Users\elsne\.aws:/root/.aws -e AWS_PROFILE=local-development -p 8080:8080 -e PORT=8080 nates-chicken-bs
```

## Giving AppRunner Permissions
This was bizzarely confusing. What wound up working was via the IAM console creating a custom [trust policy](https://stackoverflow.com/a/70092312/1333854):

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "tasks.apprunner.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

This creates a role which allows an AppRunner instance to assume this role (`sts:AssumeRole`), which would then grant it any permissions which I associate to this role (named `nates-chicken-bs-apprunner`).

Once I created this role I gave it the pre-defined `AmazonDynamoDBFullAccess` role. Then I assigned the `nates-chicken-bs-apprunner` role to the AppRunner instance and it was able to communicate with DynamoDb.

Doing this manually sucks. Since terraform is overkill for now just need to figure out the best way to automate all this setup with the CLI or something. 