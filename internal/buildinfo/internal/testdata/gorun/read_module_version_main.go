package main

import (
	"os"

	"github.com/razcoen/cqlc/internal/buildinfo"
)

func main() {
	version, err := buildinfo.ReadModuleVersion()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	_, _ = os.Stdout.WriteString(version)
}
