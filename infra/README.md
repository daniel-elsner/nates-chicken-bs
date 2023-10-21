# CDK Cheat Sheet

Check diff between code & deployed stack

```
cdk diff --profile=local-cdk
```

Generate CloudFormation template

```
cdk synth --profile=local-cdk
```

Deploy stack

```
cdk deploy --profile=local-cdk
```

Destroy stack

```
cdk destroy --profile=local-cdk
```

List stacks

```
cdk ls --profile=local-cdk
```

Running the cdk-check-diff for dev:

```
.\cdk-check-diff.bat 375842827666 us-east-2 local-cdk
```

Running the cdk-deploy-to for dev:

```
.\cdk-deploy-to.bat 375842827666 us-east-2 local-cdk
```