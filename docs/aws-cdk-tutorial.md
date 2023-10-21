# AWS CDK Setup/Tutorial
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

Then I started following this tutorial - https://docs.aws.amazon.com/cdk/v2/guide/hello_world.html

I immediately started running into issues within VSCode by creating a new CDK project within the app directory. It kept not compiling correctly
within the IDE, and the `cdk ls` command would time out.

Within the IDE it suggested creating a `go.work` file at the root, which worked, but after reading up on things that seemed like a bad approach. The idea
of a `go.work` file is that it essentially couples the two modules together, which is not what I want. The CDK is a separately managed module and should 
be able to be updated independently of the app.

What works is to open the `src` and `infrastructure` directories as separate projects in VSCode. It's kind of dumb to have to do this, but it at least works
and the `cdk ls` command no longer times out either.

Eventually I follow the tutorial some more and want to create an s3 bucket (by running `cdk deploy`), but I get this error:

```
❌ Deployment failed: Error: HelloCdkStack: This CDK deployment requires bootstrap stack version '6', but during the confirmation via SSM parameter /cdk-bootstrap/hnb659fds/version the following error occurred: AccessDeniedException: User: arn:aws:iam::375842827666:user/local-daniel-elsner is not authorized to perform: ssm:GetParameter on resource: arn:aws:ssm:us-east-2:375842827666:parameter/cdk-bootstrap/hnb659fds/version because no identity-based policy allows the ssm:GetParameter action
```

This is because I'm running the `local-development` profile, which doesn't have the permissions to create the SSM parameter. 

What I did was create a `cdk-users` user group, granting admin privilges, in AWS and create a `cdk-daniel-elsner` user. Not a great solution, but I'm just trying to 
partition permissions among profiles at least, and I imagine I'm going to need an irritating number of permissions for creating resources of any/all kinds.

Now I was able to run:

```
cdk deploy --profile=local-cdk 
```

And I get a different error:

```
 ❌ Deployment failed: Error: HelloCdkStack: SSM parameter /cdk-bootstrap/hnb659fds/version not found. Has the environment been bootstrapped? Please run 'cdk bootstrap' (see https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html)
    at Deployments.validateBootstrapStackVersion (C:\ProgramData\nvm\v16.18.0\node_modules\aws-cdk\lib\index.js:470:12023)
    at process.processTicksAndRejections (node:internal/process/task_queues:95:5)
    at async Deployments.buildSingleAsset (C:\ProgramData\nvm\v16.18.0\node_modules\aws-cdk\lib\index.js:470:10788)
    at async Object.buildAsset (C:\ProgramData\nvm\v16.18.0\node_modules\aws-cdk\lib\index.js:470:177861)
    at async C:\ProgramData\nvm\v16.18.0\node_modules\aws-cdk\lib\index.js:470:163529

HelloCdkStack: SSM parameter /cdk-bootstrap/hnb659fds/version not found. Has the environment been bootstrapped? Please run 'cdk bootstrap' (see https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html)
```

Looks more straightforward, think I just need to run `cdk bootstrap --profile=local-cdk`.

That worked, now running `cdk deploy --profile=local-cdk` again works and I have my s3 bucket.

Anyways, now I can finish the tutorial. I wind up destroying everything with `cdk destroy --profile=local-cdk`, and now I'm going to move on to setting up our actual infrastructure.