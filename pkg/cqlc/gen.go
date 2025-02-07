package cqlc

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/razcoen/cqlc/pkg/cqlc/log"
)

func Generate(config *config.Config, opts ...Option) error {
	gen, err := newGenerator(opts...)
	if err != nil {
		return fmt.Errorf("creating generator: %w", err)
	}
	return gen.Generate(config)
}

type generator struct {
	logger     log.Logger
	configPath string
}

func newGenerator(opts ...Option) (*generator, error) {
	var gen generator
	for _, opt := range opts {
		opt.apply(&gen)
	}
	if gen.logger == nil {
		slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelError,
		}))
		gen.logger = log.NewSlogAdapter(slogLogger).With("context", "codegen")
	}
	return &gen, nil
}

func (g *generator) Generate(config *config.Config) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}
	for _, config := range config.CQL {
		if config.Gen.Go == nil {
			return fmt.Errorf("golang generation config is required: only golang support")
		}
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

		logger := g.logger.With("language", "golang")
		goGenerator, err := golang.NewGenerator(logger)
		if err != nil {
			return fmt.Errorf("new go generator: %w", err)
		}
		version, err := buildinfo.ReadModuleVersion()
		if err != nil {
			logger.With("error", err).Warn("cannot evaluate the module version")
		}
		if err := goGenerator.Generate(&sdk.Context{
			Schema:      schema,
			Queries:     queries,
			SchemaPath:  config.Schema,
			QueriesPath: config.Queries,
			ConfigPath:  g.configPath,
			Version:     version,
		}, config.Gen.Go); err != nil {
			return fmt.Errorf("generate go: %w", err)
		}

	}
	return nil
}
