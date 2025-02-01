package cqlc

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
)

func Generate(config *config.Config) error {
	gen, err := newGenerator()
	if err != nil {
		return fmt.Errorf("creating generator: %w", err)
	}
	return gen.Generate(config)
}

type generator struct {
	goGenerator *golang.Generator
}

func newGenerator() (*generator, error) {
	// TODO: Logger configuration
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))
	goGenerator, err := golang.NewGenerator(logger)
	if err != nil {
		return nil, fmt.Errorf("new go generator: %w", err)
	}
	return &generator{goGenerator: goGenerator}, nil
}

func (g *generator) Generate(config *config.Config) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}
	for _, config := range config.CQL {
		sb, err := os.ReadFile(config.Schema)
		if err != nil {
			return fmt.Errorf("read schema file: %w", err)
		}
		qb, err := os.ReadFile(config.Queries)
		if err != nil {
			return fmt.Errorf("read queries file: %w", err)
		}
		sp := compiler.NewSchemaParser()
		qp := compiler.NewQueriesParser()
		schema, err := sp.Parse(string(sb))
		if err != nil {
			return fmt.Errorf("parse schema: %w", err)
		}
		queries, err := qp.Parse(string(qb))
		if err != nil {
			return fmt.Errorf("parse queries: %w", err)
		}

		if config.Gen.Overwrite {
			if err := os.RemoveAll(config.Gen.Go.Out); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("remove generated go directory: %w", err)
			}
		}

		if err := g.goGenerator.Generate(&sdk.GenerateRequest{Schema: schema, Queries: queries}, config.Gen.Go); err != nil {
			return fmt.Errorf("generate go: %w", err)
		}

	}
	return nil
}
