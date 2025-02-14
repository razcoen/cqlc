// Code generated by cqlc. DO NOT EDIT.
// cqlc version: (devel)
// schema: ../../testdata/alteringmigrations
// queries: ../../testdata/altering_migrations_queries.cql

package alteringmigrations

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/gocqlc"
)

type CreateUserParams struct {
	UserID    gocql.UUID
	Username  string
	Email     string
	LastLogin time.Time
}

func (c *Client) CreateUser(ctx context.Context, params *CreateUserParams, opts ...gocqlc.QueryOption) error {
	session := c.Session()
	q := session.Query("INSERT INTO users (user_id, username, email, last_login) VALUES (?, ?, ?, ?);", params.UserID, params.Username, params.Email, params.LastLogin)
	q = q.WithContext(ctx)
	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	if err := q.Exec(); err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	return nil
}

type FindUserParams struct {
	UserID gocql.UUID
}

type FindUserResult struct {
	UserID    gocql.UUID
	Email     string
	LastLogin time.Time
	Username  string
}

func (c *Client) FindUser(ctx context.Context, params *FindUserParams, opts ...gocqlc.QueryOption) (*FindUserResult, error) {
	session := c.Session()
	q := session.Query("SELECT * FROM users WHERE user_id = ? LIMIT 1;", params.UserID)
	q = q.WithContext(ctx)
	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	var result FindUserResult
	if err := q.Scan(&result.UserID, &result.Email, &result.LastLogin, &result.Username); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &result, nil
}
