//go:build exclude
// +build exclude

package main

import "github.com/razcoen/cqlc/pkg/cqlc"

func main() {
	if err := cqlc.Generate(&cqlc.CQLConfig{
		Queries: "./queries.cql",
		Schema:  "./schema.cql",
		Gen: &cqlc.CQLGenConfig{
			Overwrite: true,
			Go: &cqlc.CQLGenGoConfig{
				Package: "example",
				Out:     "./example",
			},
		},
	}); err != nil {
		panic(err)
	}
}
