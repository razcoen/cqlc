package cqlc

import (
	"fmt"
	"github.com/razcoen/cqlc/pkg/strfmt"
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
	out := config.Gen.Go.Out
	if err := os.Mkdir(out, 0777); err != nil && !os.IsExist(err) {
		return fmt.Errorf("create output directory: %w", err)
	}
	fn := filepath.Join(out, "client.go")
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("open client file: %w", err)
	}
	defer f.Close()
	if err := g.goGenerator.generateClient(&generateClientRequest{
		packageName: config.Gen.Go.Package,
		out:         f,
	}); err != nil {
		return fmt.Errorf("generate client: %w", err)
	}
	for _, k := range schema.Keyspaces {
		fn := filepath.Join(out, "keyspace_structs_"+strfmt.ToSingularSnakeCase(k.Name)+".go")
		f, err := os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("open keyspace file: %w", err)
		}
		// TODO: Abstract Go Generator by having a single method: Generate(schema, queries)
		// TODO: Schema struct redundant?
		defer f.Close()
		resp, err := g.goGenerator.generateKeyspaceStructs(&generateKeyspaceStructsRequest{
			keyspace:    k,
			packageName: config.Gen.Go.Package,
			out:         f,
		})
		if err != nil {
			return fmt.Errorf("generate keyspace: %w", err)
		}
		fn = filepath.Join(out, "keyspace_queries_"+strfmt.ToSingularSnakeCase(k.Name)+".go")
		f, err = os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("open queries file: %w", err)
		}
		defer f.Close()
		// TODO: Aggregate across keyspaces
		if err := g.goGenerator.generateQueries(&generateQueriesRequest{
			queries:           queries,
			structByTableName: resp.structByTableName,
			packageName:       config.Gen.Go.Package,
			out:               f,
		}); err != nil {
			return fmt.Errorf("generate queries: %w", err)
		}
	}
	return nil
}
