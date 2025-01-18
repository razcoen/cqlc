package gocqlc

import "github.com/gocql/gocql"

type QueryOption interface {
	Apply(query *gocql.Query) *gocql.Query
}

type queryOptionFunc func(query *gocql.Query) *gocql.Query

func (f queryOptionFunc) Apply(query *gocql.Query) *gocql.Query { return f(query) }

func WithPageState(state []byte) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.PageState(state)
	})
}

func WithPageSize(n int) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.PageSize(n)
	})
}
