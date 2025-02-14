package gocqlc

import "github.com/gocql/gocql"

type QueryOption interface {
	Apply(query *gocql.Query) *gocql.Query
}

type QueryOptionFunc func(query *gocql.Query) *gocql.Query

func (f QueryOptionFunc) Apply(query *gocql.Query) *gocql.Query { return f(query) }

type BatchOption interface {
	Apply(*gocql.Batch) *gocql.Batch
}

type BatchOptionFunc func(*gocql.Batch) *gocql.Batch

func (f BatchOptionFunc) Apply(batch *gocql.Batch) *gocql.Batch { return f(batch) }

func WithBatchType(batchType gocql.BatchType) BatchOption {
	return BatchOptionFunc(func(batch *gocql.Batch) *gocql.Batch {
		batch.Type = batchType
		return batch
	})
}
