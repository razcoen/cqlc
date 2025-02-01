package config

import (
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"gopkg.in/yaml.v3"
)

type Config struct {
	CQL []*CQL `yaml:"cql" validate:"required,min=1,dive"`
}

type CQL struct {
	Queries string  `yaml:"queries" validate:"required"`
	Schema  string  `yaml:"schema" validate:"required"`
	Gen     *CQLGen `yaml:"gen" validate:"required"`
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

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
