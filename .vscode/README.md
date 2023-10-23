# .vscode

Checking VSCode specific things in here. Not a requirement to use VS Code, but it 
seems to work fine for Go development.

## Extensions

### REST Client
https://marketplace.visualstudio.com/items?itemName=humao.rest-client

This is what allows me to run the `.http` files. It's nice for testing out API 
calls locally.

## launch.json
This is the configuration for debugging. It's assuming that you've ran 
`aws configure --profile local-development` at some point to create an aws configuration called local-development, which is why it's setting that for the environment variable.