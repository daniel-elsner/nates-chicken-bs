# DynamoDb Emulation

Junk notes for how I set this up:

```
docker pull amazon/dynamodb-local
```

```
docker run -p 8000:8000 amazon/dynamodb-local
```

Should have it running on localhost:8000 at this point. Verify with something like:

```
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

Now I need to figure out how to point the app to this local dynamo location.

Seems like you do this by hardcoding `endpoint` on the `aws.Config` object:

```go
func InitDynamoClient() {
	config := config.GetAWSConfig()

	endpoint := "http://localhost:8000"
	config.Endpoint = &endpoint

	sess := session.Must(session.NewSession(config))
	DynamoDBClient = dynamodb.New(sess)
}
```

Obviously this is gross, we need it to be configurable via some kind of appsettings or environment variables. I'm not sure how these kinds of configs are most commonly dealt with in go, so that's something to figure out separately. For now just going to leave this commented out so it's easy to toggle.