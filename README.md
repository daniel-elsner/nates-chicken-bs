# nates-chicken-bs

This repository contains a simple Go application which is deployed to AWS AppRunner. It presents a basic REST API which could eventually turn into a backend for Nate's Chicken BS mobile application.

Infrastructure is managed with AWS CDK in the [ncbs-cdk](https://github.com/daniel-elsner/ncbs-cdk) repository.

# Running the App

## Prerequisites

To run this you will need:
 - [Go](https://go.dev/doc/install)

Nice-To-Haves:
 - [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
 - [Docker](https://docs.docker.com/get-docker/)

## Docker Commands
This will build the image and run it on port 8080 with the environment variable PORT set to 8080. This is on Docker for Windows, should be similar on other platforms.

```
docker build -t nates-chicken-bs .
```

```
docker run -p 8080:8080 -e PORT=8080 nates-chicken-bs
```

This command will mount the AWS credentials from (my) host machine into the container. This is useful for running the app locally in docker (steps over on [infrastructure-notes](docs/infrastructure-notes.md) for I set up the AWS CLI on my machine).

```
docker run -v ${env:USERPROFILE}\.aws:/root/.aws -p 8080:8080 -e PORT=8080 nates-chicken-bs
```

This gets even bigger if you don't want to use a `[default]` profile in your credentials file. You can specify the profile name with the `AWS_PROFILE` environment variable. This is useful for running the app locally in Docker.

```
docker run -v ${env:USERPROFILE}\.aws:/root/.aws -e AWS_PROFILE=local-development -p 8080:8080 -e PORT=8080 nates-chicken-bs
```


## General Commands

To build the package:
```
go build 
```

Then simply run the executable: 
```
.\ncbs.exe
```

# Setting up Air for Live Reload
https://github.com/cosmtrek/air

First, install air

```bash
go install github.com/cosmtrek/air@latest
```

Verify the installation
```bash
air -v
```

Run air in the root of the project
```bash
air
```