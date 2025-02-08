package gocqlc

import "github.com/gocql/gocql"

type QueryOption interface {
	Apply(query *gocql.Query) *gocql.Query
}

type QueryOptionFunc func(query *gocql.Query) *gocql.Query

func (f QueryOptionFunc) Apply(query *gocql.Query) *gocql.Query { return f(query) }

type BatchOption interface {
	Apply(query *gocql.Batch) *gocql.Batch
}

type BatchOptionFunc func(query *gocql.Batch) *gocql.Batch

func (f BatchOptionFunc) Apply(query *gocql.Batch) *gocql.Batch { return f(query) }
