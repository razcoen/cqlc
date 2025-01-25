//go:build exclude
// +build exclude

package main

import (
	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
)

func main() {
	if err := cqlc.Generate(&config.Config{
		CQL: []*config.CQL{
			{
				Queries: "queries.cql",
				Schema:  "schema.cql",
				Gen: &config.CQLGen{
					Overwrite: true,
					Go: &golang.Options{
						Package: "example",
						Out:     "example",
					},
				},
			},
		},
	}); err != nil {
		panic(err)
	}
}
