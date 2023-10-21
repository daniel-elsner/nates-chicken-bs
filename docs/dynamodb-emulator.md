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