package config

import (
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"gopkg.in/yaml.v3"
)

type Config struct {
	// CQL is a slice of CQL configurations.
	// It is required and must contain at least one element.
	CQL []*CQL `yaml:"cql" validate:"required,min=1,dive"`
}

type CQL struct {
	// Queries is a filepath containing the CQL queries.
	Queries string `yaml:"queries" validate:"required"`
	// Schema is a filepath containing the schema definition.
	Schema string `yaml:"schema" validate:"required"`
	// Gen is a pointer to a CQLGen struct containing generation options.
	Gen *CQLGen `yaml:"gen" validate:"required"`
}

type CQLGen struct {
	// Overwrite is a boolean indicating whether to overwrite existing files.
	Overwrite bool `yaml:"overwrite"`
	// Go is a pointer to golang.Options containing Go-specific generation options.
	Go *golang.Options `yaml:"go"`
}

// ParseConfig reads and parses the configuration from the provided io.Reader.
// It expects the configuration to be in YAML format and validates the parsed
// configuration using the validator package. If the configuration is invalid
// or cannot be decoded, an error is returned.
func ParseConfig(r io.Reader) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		return nil, fmt.Errorf("decode config file: %w", err)
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}
	return &config, config.Validate()
}

// Validate checks the Config struct to ensure it meets the required
// validation rules.
func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
