# nates-chicken-bs

This repository contains a simple Go application which is deployed to AWS AppRunner. It presents a basic REST API which could eventually turn into a backend for Nate's Chicken BS mobile application.

Infrastructure is managed with AWS CDK in the [ncbs-cdk](https://github.com/daniel-elsner/ncbs-cdk) repository.

This is built & deployed on merges to `main` using GitHub Actions. The workflow is defined [here](.github/workflows/build-and-push-image.yml).

# Running the App

## Prerequisites

To run this you should have:
 - [Go](https://go.dev/doc/install)
 - [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
 - [Docker](https://docs.docker.com/get-docker/)

## Configure Profile

Configure the AWS CLI with a profile named `local-development`. The profile name is important as it is assumed to exist in various configurations & scripts.

```bash
aws configure --profile local-development
```

You will need to be given the access key and secret key for the profile. These can be generated in the AWS console (we should see if SSO is viable at some point).

## Run the App

You have a few options for running locally. You can: 

1. Build the project and run the executable directly
2. Launch it in Debug mode from VS Code
3. Run it in Docker


### Option 1: Build and Run

Navigate to the `src` directory and execute the `run-build-local.bat` script:

```
cd src
```

```
\run-build-local.bat
```

### Option 2: Debug in VS Code

For this you should be able to leverage the `launch.json` file in the `.vscode` directory at the root of this project.

Simply hit F5 to launch the debugger, and it should use the `ncbs - Debug Mode` configuration.

### Option 3: Run in Docker

Navigate to the `src` directory and execute the `run-docker.bat` script:

```
cd src
```

```
\run-docker.bat
```

## Emulating Dependencies

By default, the app will run against resources in AWS. Managing permissions for this can be difficult as we add more dependencies, so we can try instead to emulate them locally.

This is what the `local-env` directory is for. It contains a `docker-compose` file which will spin up local versions of the dependencies (such as DynamoDB) which the app can then connect to. 

To leverage this, navigate to the `local-env` directory and execute the `setup-local-env.bat` script:

```
cd local-env
```

```
\setup-local-env.bat
```

**NOTE:** This presently works for setting up the dependencies, but I haven't cleanly updated the app to connect to them. I'll do that soon.

## Setting up Air for Live Reload

[air](https://github.com/cosmtrek/air) is a nice tool for local development which will automatically reload the app when it detects changes to the source code.

First, install air:

```bash
go install github.com/cosmtrek/air@latest
```

Verify the installation

```bash
air -v
```

Run air from the `/src` directory:

```bash
cd /src
air
```