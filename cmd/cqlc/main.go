package main

import (
	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/razcoen/cqlc/internal/cqlc"
)

func main() {
	if err := storeBuildInfo(); err != nil {
		// TODO: Should panic?
		panic(err)
	}
	if err := cqlc.Run(); err != nil {
		// TODO: Should panic?
		panic(err)
	}
}

var version = "v0.0.0-dev"

func storeBuildInfo() error {
	return buildinfo.Store(&buildinfo.Flags{
		Version: version,
	})
}
