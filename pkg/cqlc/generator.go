package cqlc

import (
	"errors"
	"fmt"
	"os"
)

func Generate(config *CQLConfig) error {
	gen, err := NewGenerator()
	if err != nil {
		return fmt.Errorf("creating generator: %w", err)
	}
	return gen.Generate(config)
}

type Generator struct {
	goGenerator *goGenerator
}

func NewGenerator() (*Generator, error) {
	goGenerator, err := newGoGenerator()
	if err != nil {
		return nil, fmt.Errorf("new go generator: %w", err)
	}
	return &Generator{goGenerator: goGenerator}, nil
}

func (g *Generator) Generate(config *CQLConfig) error {
	sb, err := os.ReadFile(config.Schema)
	if err != nil {
		return fmt.Errorf("read schema file: %w", err)
	}
	qb, err := os.ReadFile(config.Queries)
	if err != nil {
		return fmt.Errorf("read queries file: %w", err)
	}
	sp := NewSchemaParser()
	qp := NewQueriesParser()
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

	if err := g.goGenerator.generate(config.Gen.Go, schema, queries); err != nil {
		return fmt.Errorf("generate go: %w", err)
	}
	return nil
}
