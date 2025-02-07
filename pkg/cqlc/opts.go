package cqlc

import (
	"github.com/razcoen/cqlc/pkg/cqlc/log"
)

type Option interface {
	apply(*generator)
}

type optionFunc func(*generator)

func (f optionFunc) apply(g *generator) { f(g) }

func WithLogger(logger log.Logger) Option {
	return optionFunc(func(g *generator) {
		g.logger = logger
	})
}

func WithConfigPath(configPath string) Option {
	return optionFunc(func(g *generator) {
		g.configPath = configPath
	})
}
