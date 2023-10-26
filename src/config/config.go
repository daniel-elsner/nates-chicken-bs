package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	DeployEnv string `envconfig:"DEPLOY_ENV"`
	Port      string `envconfig:"PORT"`
}

// LoadConfiguration loads the configuration from environment variables. This should
// be the only place where environment variables are read with the exception of things
// like AWS_PROFILE since those are not application specific.
func LoadConfiguration() (*Configuration, error) {
	var config Configuration
	if err := envconfig.Process("ncbs", &config); err != nil {
		return &config, err
	}

	if config.DeployEnv == "" {
		return &config, errors.New("DEPLOY_ENV must be set")
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return &config, nil
}
