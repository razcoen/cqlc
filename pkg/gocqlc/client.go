package gocqlc

import (
	"errors"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/log"
)

var ErrNilSession = errors.New("session cannot be nil")
var ErrClosedSession = errors.New("session cannot be closed")

type Client struct {
	logger              log.Logger
	session             *gocql.Session
	defaultQueryOptions []QueryOption
	defaultBatchOptions []BatchOption
}

func NewClient(session *gocql.Session, opts ...ClientOption) (*Client, error) {
	if session == nil {
		return nil, ErrNilSession
	}
	if session.Closed() {
		return nil, ErrClosedSession
	}
	client := &Client{
		session: session,
		logger:  log.NopLogger(),
	}
	for _, opt := range opts {
		opt.apply(client)
	}
	return client, nil
}

func (c *Client) Logger() log.Logger                 { return c.logger }
func (c *Client) Session() *gocql.Session            { return c.session }
func (c *Client) DefaultQueryOptions() []QueryOption { return c.defaultQueryOptions }
func (c *Client) DefaultBatchOptions() []BatchOption { return c.defaultBatchOptions }

type ClientOption interface {
	apply(*Client)
}

var _ ClientOption = clientOptionFunc(nil)

type clientOptionFunc func(*Client)

func (f clientOptionFunc) apply(c *Client) {
	f(c)
}

// WithLogger sets the logger for the client.
func WithLogger(logger log.Logger) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.logger = logger
	})
}

// WithDefaultQueryOptions sets the default query options for the client.
// The default query options will be applied prior to the additional options.
func WithDefaultQueryOptions(opts ...QueryOption) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.defaultQueryOptions = opts
	})
}

// WithDefaultBatchOptions sets the default batch options for the client.
// The default batch options will be applied prior to the additional options.
func WithDefaultBatchOptions(opts ...BatchOption) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.defaultBatchOptions = opts
	})
}
