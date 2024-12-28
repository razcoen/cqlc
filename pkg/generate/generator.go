package generate

import (
	"fmt"
	"os"

	"github.com/razcoen/cqlc/pkg/cql"
	cqlcv1 "github.com/razcoen/cqlc/pkg/proto/clqc/v1"
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

func (g *Generator) Generate(config *cqlcv1.GenerateConfig) error {
	b, err := os.ReadFile(config.Schema)
	if err != nil {
		return fmt.Errorf("read schema file: %w", err)
	}
	sp := cql.NewSchemaParser()
	schema, err := sp.Parse(string(b))
	if err != nil {
		return fmt.Errorf("parse schema: %w", err)
	}
	for _, k := range schema.Keyspaces {
		g.goGenerator.generateKeyspace(&generateKeyspaceRequest{
			keyspace:    k,
			packageName: config.Go.Package,
      out:         nil, // TODO: Create the files and etc.
		})
	}
	return nil
}
