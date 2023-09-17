package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicConfig  PublicConfig
	PrivateConfig PrivateConfig
}

type PublicConfig struct {
	PublicKey string
}
type PrivateConfig struct {
	SecretKey string
}

var (
	envs = map[string]PublicConfig{
		"local": {
			PublicKey: "local",
		},
		"dev": {
			PublicKey: "dev",
		},
		"prod": {
			PublicKey: "prod",
		},
	}
)

func NewConfig(env string) (Config, error) {
	publicConfig, exists := envs[env]
	if !exists {
		return Config{}, errors.New("env doest not exists")
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return Config{}, errors.New("Secret key doest not exists in env.")
	}
	return Config{
		PublicConfig: publicConfig,
		PrivateConfig: PrivateConfig{
			SecretKey: secretKey,
		},
	}, nil
}
