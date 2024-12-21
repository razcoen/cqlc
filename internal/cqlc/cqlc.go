package cqlc

import "github.com/razcoen/cqlc/internal/cqlc/cmd"

func Run() error {
	return cmd.NewRootCommand().Execute()
}
