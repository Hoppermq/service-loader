package config

import (
	"embed"
	"log/slog"

	"github.com/knadh/koanf/v2"
)

// SecretLoader is an interface that can be implemented by the target to load secrets from SSM.
type SecretLoader interface {
	// Load loads a secret from the secret manager (e.g. aws ssm).
	// The lookup secret path will be constructed as follows:
	// /rancher2/{env}/{stack}/{serviceName}/{secretPath}
	Load(intoKey string, secretPath string)
}

type ConfigSecretLoader interface {
	SecretLoader
	SetConfigLoader(configLoader *koanf.Koanf)
	SetLogger(logger *slog.Logger)
}

type config struct {
  target  interface{}
  configLoader *koanf.Koanf
  fname  string
  fs *embed.FS
  secretLoader SecretLoader
}

func (c *config)loader() error {
  return nil
}
