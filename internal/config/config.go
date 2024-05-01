package config

import (
	"fmt"
	"os"

	"github.com/sabahtalateh/di"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB DB `yaml:"db"`
}

type DB struct {
	DSN string `yaml:"dsn"`
}

func readConfig() (*Config, error) {
	confPath := os.Getenv("CONF_PATH")
	if confPath == "" {
		return nil, fmt.Errorf("CONF_PATH env var not set")
	}

	confBytes, err := os.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	conf := new(Config)
	if err = yaml.Unmarshal(confBytes, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func SetupConfig(c *di.Container) error {
	return di.Setup[*Config](c,
		di.InitE(func(c *di.Container) (*Config, error) { return readConfig() }),
	)
}
