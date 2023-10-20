# AWS CDK
Dumping ground for notes as I work through this, as I've never used the AWS CDK before.

## Steps

First, install the CDK with `npm`:

```
npm install -g aws-cdk
```

Apparently I was on an old version of node, so it gave me a warning about that and I upgraded (from here https://nodejs.org/en/download).

To verify installation:
```
cdk --version
```

Then I created a new project in the `/infrastructure` directory:

```
cd infrastructure
cdk init app --language=go
```

Then I downloaded various go packages I figured I would need:

```
go get github.com/aws/aws-cdk-go/awscdk/v2
go get github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner
go get github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb
go get github.com/aws/aws-cdk-go/awscdk/v2/awsecr
go get github.com/aws/aws-cdk-go/awscdk/v2/awsiam
```