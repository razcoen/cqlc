package cqlc

import (
	"fmt"
	"github.com/razcoen/cqlc/pkg/cqlc/strfmt"
	"os"
	"path/filepath"
)

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
	b, err := os.ReadFile(config.Schema)
	if err != nil {
		return fmt.Errorf("read schema file: %w", err)
	}
	sp := NewSchemaParser()
	schema, err := sp.Parse(string(b))
	if err != nil {
		return fmt.Errorf("parse schema: %w", err)
	}
	out := config.Gen.Go.Out
	if err := os.Mkdir(out, 0777); err != nil && !os.IsExist(err) {
		return fmt.Errorf("create output directory: %w", err)
	}
	for _, k := range schema.Keyspaces {
		fn := filepath.Join(out, "keyspace_"+strfmt.ToSingularSnakeCase(k.Name)+".go")
		f, err := os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("open keyspace file: %w", err)
		}
		if err := g.goGenerator.generateKeyspace(&generateKeyspaceRequest{
			keyspace:    k,
			packageName: config.Gen.Go.Package,
			out:         f,
		}); err != nil {
			return fmt.Errorf("generate keyspace: %w", err)
		}
	}
	return nil
}
