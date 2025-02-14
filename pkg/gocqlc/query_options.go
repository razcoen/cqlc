package gocqlc

import "github.com/gocql/gocql"

// QueryOption represents an option to modify for a query.
type QueryOption interface {
	Apply(query *gocql.Query) *gocql.Query
}

var _ QueryOption = queryOptionFunc(nil)

type queryOptionFunc func(query *gocql.Query) *gocql.Query

func (f queryOptionFunc) Apply(query *gocql.Query) *gocql.Query { return f(query) }

// BatchOption represents an option to modify for a batch query.
type BatchOption interface {
	Apply(*gocql.Batch) *gocql.Batch
}

var _ BatchOption = batchOptionFunc(nil)

type batchOptionFunc func(*gocql.Batch) *gocql.Batch

func (f batchOptionFunc) Apply(batch *gocql.Batch) *gocql.Batch { return f(batch) }

// WithBatchType returns a BatchOption that sets the batch type.
func WithBatchType(batchType gocql.BatchType) BatchOption {
	return batchOptionFunc(func(batch *gocql.Batch) *gocql.Batch {
		batch.Type = batchType
		return batch
	})
}
