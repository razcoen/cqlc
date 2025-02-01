package main

import (
	"os"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/razcoen/cqlc/internal/cqlc"
)

var version = "v0.0.0-dev"

func main() {
	if err := cqlc.Run(cqlc.WithBuildFlags(&buildinfo.Flags{Version: version})); err != nil {
		os.Exit(1)
	}
}
