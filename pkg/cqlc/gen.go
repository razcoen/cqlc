package cqlc

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/razcoen/cqlc/pkg/log"
)

func Generate(config *config.Config, opts ...Option) error {
	gen := newGenerator(opts...)
	return gen.Generate(config)
}

type generator struct {
	logger     log.Logger
	configPath string
}

func newGenerator(opts ...Option) *generator {
	var gen generator
	for _, opt := range opts {
		opt.apply(&gen)
	}
	if gen.logger == nil {
		slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelWarn}))
		gen.logger = log.NewSlogAdapter(slogLogger).With("context", "codegen")
	}
	return &gen
}

func (gen *generator) Generate(config *config.Config) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}
	for _, config := range config.CQL {
		if config.Gen.Go == nil {
			return fmt.Errorf("golang generation config is required: only golang support")
		}
		logger := gen.logger.
			With("schema", config.Schema).
			With("queries", config.Queries)
		logger.Debug("test123")
		f, err := os.Stat(config.Schema)
		if err != nil {
			return fmt.Errorf("stat schema file: %w", err)
		}
		var schemaContents string
		if f.IsDir() {
			// If the given schema file is indeed a directory, then assuming this is a migrations directory.
			// A migrations directory is assumed to have the following attributes:
			// 	1. Migrations directory is a flat list of files.
			// 	2. Files with the keyword "down" are considered down migrations.
			// 	3. Migrations should be run in an alphabetic ascending order.
			// Therefore, given these assumptions, flatten the migrations into a single schema file.
			// TODO: Need to test this by supporting alter table.
			downMigrationKeywords := []string{"down"}
			logger.Info("assuming provided schema is a migrations directory: provided schema path is a directory")
			entries, err := os.ReadDir(config.Schema)
			if err != nil {
				return fmt.Errorf("read schema migrations directory: %w", err)
			}
			var migrations []string
			for _, e := range entries {
				file := filepath.Join(config.Schema, e.Name())
				logger := logger.With("file", file)
				if e.IsDir() {
					logger.Debug("skipping file: only flat migrations structure is supported")
					continue
				}
				for _, keyword := range downMigrationKeywords {
					logger := logger.With("keyword", keyword)
					if strings.Contains(e.Name(), keyword) {
						// TODO: Provide an option to override the keywords.
						logger.Debug(`skipping file: file contains a down migration keyword`)
						continue
					}
					migrations = append(migrations, file)
				}
			}
			// Sort the migrations in ascending order, assuming that they are ordered alphabetically.
			slices.Sort(migrations)
			logger.With("migrations", migrations).Debug("assuming migrations are orderd alphabetically")
			var cql strings.Builder
			for _, migration := range migrations {
				logger := logger.With("file", migration)
				b, err := os.ReadFile(migration)
				if err != nil {
					logger.With("error", err).Error("failed to read migration file")
					return fmt.Errorf("read migration file: %w", err)
				}
				_, _ = cql.Write(b)
				_, _ = cql.WriteString("\n") // for formatting purposes
			}
			schemaContents = cql.String()
		} else {
			sb, err := os.ReadFile(config.Schema)
			if err != nil {
				return fmt.Errorf("read schema file: %w", err)
			}
			schemaContents = string(sb)
		}
		qb, err := os.ReadFile(config.Queries)
		if err != nil {
			return fmt.Errorf("read queries file: %w", err)
		}
		sp := compiler.NewSchemaParser()
		// TODO: Since migrations are flattened, errors from the parsing are shown in incorrect lines.
		// TODO: Run some parsing logic, without actually building the schema, to allow better error handling by the user.
		logger.
			With("schema.contents", schemaContents).
			Debug("parsing the following schema")
		schema, err := sp.Parse(schemaContents)
		if err != nil {
			return fmt.Errorf("parse schema: %w", err)
		}
		qp := compiler.NewQueriesParser()
		queries, err := qp.Parse(string(qb))
		if err != nil {
			return fmt.Errorf("parse queries: %w", err)
		}

		if config.Gen.Overwrite {
			if err := os.RemoveAll(config.Gen.Go.Out); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("remove generated go directory: %w", err)
			}
		}

		goGenerator, err := golang.NewGenerator(logger.With("language", "golang"))
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
			ConfigPath:  gen.configPath,
			Version:     version,
		}, config.Gen.Go); err != nil {
			return fmt.Errorf("generate go: %w", err)
		}

	}
	return nil
}
