package cqlc

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Version string       `yaml:"version"`
	CQL     []*CQLConfig `yaml:"cql"`
}

type CQLConfig struct {
	Queries string        `yaml:"queries"`
	Schema  string        `yaml:"schema"`
	Gen     *CQLGenConfig `yaml:"gen"`
}

type CQLGenConfig struct {
	Overwrite bool            `yaml:"overwrite"`
	Go        *CQLGenGoConfig `yaml:"go"`
}

type CQLGenGoConfig struct {
	Package string `yaml:"package"`
	Out     string `yaml:"out"`
}

func ReadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file: %w", err)
	}
	var config Config
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, fmt.Errorf("decode config file: %w", err)
	}
	if config.Version != "1" {
		return nil, fmt.Errorf("unsupported config version: %s", config.Version)
	}
	return &config, nil
}
