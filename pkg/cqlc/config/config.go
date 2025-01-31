package config

import (
	"fmt"
	"io"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"gopkg.in/yaml.v3"
)

type Config struct {
	CQL []*CQL `yaml:"cql"`
}

type CQL struct {
	Queries string  `yaml:"queries"`
	Schema  string  `yaml:"schema"`
	Gen     *CQLGen `yaml:"gen"`
}

type CQLGen struct {
	Overwrite bool            `yaml:"overwrite"`
	Go        *golang.Options `yaml:"go"`
}

func ParseConfig(r io.Reader) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		return nil, fmt.Errorf("decode config file: %w", err)
	}
	return &config, nil
}
