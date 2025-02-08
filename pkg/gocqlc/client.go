package gocqlc

import (
	"errors"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/log"
)

var ErrNilSession = errors.New("session cannot be nil")
var ErrClosedSession = errors.New("session cannot be closed")

type Client struct {
	logger  log.Logger
	session *gocql.Session
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

func (c *Client) Logger() log.Logger      { return c.logger }
func (c *Client) Session() *gocql.Session { return c.session }

type ClientOption interface {
	apply(*Client)
}

var _ ClientOption = clientOptionFunc(nil)

type clientOptionFunc func(*Client)

func (f clientOptionFunc) apply(c *Client) {
	f(c)
}

func WithLogger(logger log.Logger) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.logger = logger
	})
}
