package gocqlc

import "github.com/gocql/gocql"

// QueryOption represents an option to modify for a query.
type QueryOption interface {
	apply(query *gocql.Query) *gocql.Query
}

var _ QueryOption = queryOptionFunc(nil)

type queryOptionFunc func(query *gocql.Query) *gocql.Query

func (f queryOptionFunc) apply(query *gocql.Query) *gocql.Query { return f(query) }

// ApplyQueryOptions applies a list of QueryOptions to a query.
func ApplyQueryOptions(query *gocql.Query, opts ...QueryOption) *gocql.Query {
	for _, opt := range opts {
		query = opt.apply(query)
	}
	return query
}

func WithConsistency(consistency gocql.Consistency) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.Consistency(consistency)
	})
}

func WithSerialConsistency(serialConsistency gocql.SerialConsistency) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.SerialConsistency(serialConsistency)
	})
}

func WithTimestamp(timestamp int64) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.WithTimestamp(timestamp)
	})
}

func WithTrace(tracer gocql.Tracer) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.Trace(tracer)
	})
}

// WithPageSize returns a QueryOption that sets the page size for a query.
func WithPageSize(pageSize int) QueryOption {
	return queryOptionFunc(func(query *gocql.Query) *gocql.Query {
		return query.PageSize(pageSize)
	})
}

// BatchOption represents an option to modify for a batch query.
type BatchOption interface {
	apply(*gocql.Batch) *gocql.Batch
}

var _ BatchOption = batchOptionFunc(nil)

type batchOptionFunc func(*gocql.Batch) *gocql.Batch

func (f batchOptionFunc) apply(batch *gocql.Batch) *gocql.Batch { return f(batch) }

func ApplyBatchOptions(batch *gocql.Batch, opts ...BatchOption) *gocql.Batch {
	for _, opt := range opts {
		batch = opt.apply(batch)
	}
	return batch
}

// WithBatchType returns a BatchOption that sets the batch type.
func WithBatchType(batchType gocql.BatchType) BatchOption {
	return batchOptionFunc(func(batch *gocql.Batch) *gocql.Batch {
		batch.Type = batchType
		return batch
	})
}
