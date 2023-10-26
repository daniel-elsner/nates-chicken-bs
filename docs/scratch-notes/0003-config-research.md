# App Config Options

## Viper

Viper seems to be a pretty common tool for doing configs in Go apps.

https://github.com/spf13/viper

The documentation makes note of supporting "12 factor applications". I don't know what that is, but it sounds like something I should know about.

I'm assuming they're talking about this - https://12factor.net/config

Meaning, doing everything via environment variables. The reasoning seems to be that you can't really accidentally check in environment variables to source control, so it's a good way to keep secrets out of your codebase and ensure you can always change them without having to recompile/redeploy. 

Either way, Viper seems sort of overkill. I don't want/need to support a dozen different config approaches. 

## Environment Variables Only?

I do like the approach of only having environment variables. Managing config files is difficult/annoying, although it is easier for local development.

When it comes to secrets, AppRunner supports injecting values from SSM/Secrets Manager as environment variables, so that means we never need to explicitly pull from those services with app code, which is nice.

## envconfig

This seems like a nice lightweight library for managing environment variables - https://github.com/kelseyhightower/envconfig

Assuming I go with all env vars + this, the question would be how to handle the config in code. 

I'm thinking a single `config.go` file which has a struct with all the config values, and then a `config.LoadConfiguration()` function which calls `envconfig.Process()` to populate the struct. I can call that function from `main()` and then pass the config struct around as needed.

## Config Approach

So if we settle on everything being environment variables + using `envconfig` to manage them then we can have something like:

### Local Development

Manage setting environment variables with scripting. We want a single definition of the environment variables, and then a way to set them all at once for local development, whether we're running the executable or in the container.

This can be managed with something like a `.env` file which is gitignored, and then we can parse that file and set the environment variables with a script when executing the `run-build-local.bat` file, or we can just hand the file to the container when running it locally with `docker run` in the `run-docker.bat` file.

### AppRunner

For AppRunner we should simply manage these via the AppRunner config (meaning the definition of the AppRunner service itself in 
the CDK repository). This is nice because there's no need to manage any deployment related configs in this code base, and there 
shouldn't be any way to screw them up/over-write them by accident (unless someone screws up the CDK config).

