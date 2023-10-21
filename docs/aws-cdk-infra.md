# AWS CDK Infra
Dumping ground of notes as I try to set up our "real" infrastructure.

## Steps
The first odd thing I ran into was how to specify an appropriate environment.

It says for dev environments to uncomment this bit of code:

```go
// Uncomment to specialize this stack for the AWS Account and Region that are
// implied by the current CLI configuration. This is recommended for dev
// stacks.
//---------------------------------------------------------------------------
return &awscdk.Environment{
    Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
    Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
}
```

This means setting environment variables, so I did:

```
$env:CDK_DEFAULT_ACCOUNT="375842827666"
$env:CDK_DEFAULT_REGION="us-east-2"
```

Then I boostrapped the project:

```
cdk bootstrap --profile local-cdk
```

But this still didn't use the us-east-2 region which we want, so I had to manually do:

```
cdk bootstrap aws://375842827666/us-east-2 --profile=local-cdk
```

And despite all this, it still kept using us-east-1, the only thing that worked was setting

```
$env:AWS_DEFAULT_REGION="us-east-2"
```

Which is beyond stupid since I'm explicitly pulling the env variables in the code, I can even log them out and see it.

Anyways, reading through this - https://docs.aws.amazon.com/cdk/v2/guide/environments.html - has some helpful suggestions.
A good approach seems to be some bash scripts to set the environment variables and then trigger the cdk commands.

```
@findstr /B /V @ %~dpnx0 > %~dpn0.ps1 && powershell -ExecutionPolicy Bypass %~dpn0.ps1 %*
@exit /B %ERRORLEVEL%
if ($args.length -ge 3) {
    $env:CDK_DEPLOY_ACCOUNT, $args = $args
    $env:CDK_DEPLOY_REGION,  $args = $args
    $profile, $args = $args
    npx cdk deploy $args --profile=$profile
    exit $lastExitCode
} else {
    [console]::error.writeline("Provide account, region, and profile as the first three args.")
    [console]::error.writeline("Additional args are passed through to cdk deploy.")
    exit 1
}
```

Created this as cdk-deploy-to.bat and it can be run like so:

```
.\cdk-deploy-to.bat 375842827666 us-east-2 local-cdk
```

This works and deploys to the correct region, but it's still stupid that it's necessary.

For some reason, after changing things in the code to use `CDK_DEPLOY_ACCOUNT` and `CDK_DEPLOY_REGION` instead of `CDK_DEFAULT_ACCOUNT` and `CDK_DEFAULT_REGION`, it's happy 
and I can just run `cdk deploy --profile=local-cdk` and it works (after setting the environment variables).

```
$env:CDK_DEPLOY_ACCOUNT="375842827666"
$env:CDK_DEPLOY_REGION="us-east-2"
```

Either way at least the bat script works, so that's a convenient thing to use in a more automated/pipeline fashion.