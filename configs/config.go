package configs

import (
	"github.com/greencoda/confiq"
	envLoader "github.com/greencoda/confiq/loaders/env"
	jsonLoader "github.com/greencoda/confiq/loaders/json"
)

type Config struct {
	PostgresUsername string `cfg:"POSTGRES_USERNAME"`
	PostgresPassword string `cfg:"POSTGRES_PASSWORD"`
	PostgresHost     string `cfg:"POSTGRES_HOST"`
	PostgresPort     string `cfg:"POSTGRES_PORT"`
	PostgresDBName   string `cfg:"POSTGRES_DATABASE_NAME"`
	Debug            bool   `cfg:"DEBUG"`
}

func New() (*Config, error) {
	container := envLoader.Load()
	container.FromEnvironment()

	configSet := confiq.New()
	configSet.Load(container)
	cfg := Config{}
	if err := configSet.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewFromString(configJson string) (*Config, error) {
	configSet := confiq.New()
	configSet.Load(jsonLoader.Load().FromString(configJson))
	cfg := Config{}
	if err := configSet.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
