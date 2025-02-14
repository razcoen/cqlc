// Code generated by cqlc. DO NOT EDIT.
// cqlc version: (devel)
// config: ../cqlc.yaml
// schema: ../schema.cql
// queries: ../queries.cql

package example

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/gocqlc"
	"github.com/razcoen/cqlc/pkg/log"
)

type CreateUserParams struct {
	UserID    gocql.UUID
	Username  string
	Email     string
	CreatedAt time.Time
}

func (c *Client) CreateUser(ctx context.Context, params *CreateUserParams, opts ...gocqlc.QueryOption) error {
	session := c.Session()
	q := session.Query("INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);", params.UserID, params.Username, params.Email, params.CreatedAt)
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

type CreateUsersParams struct {
	UserID    gocql.UUID
	Username  string
	Email     string
	CreatedAt time.Time
}

func (c *Client) CreateUsers(ctx context.Context, params []*CreateUsersParams, opts ...gocqlc.BatchOption) error {
	session := c.Session()
	b := session.NewBatch(gocql.LoggedBatch)
	for _, v := range params {
		b.Query("INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);", v.UserID, v.Username, v.Email, v.CreatedAt)
	}
	b = b.WithContext(ctx)
	for _, opt := range c.DefaultBatchOptions() {
		b = opt.Apply(b)
	}
	for _, opt := range opts {
		b = opt.Apply(b)
	}
	if err := session.ExecuteBatch(b); err != nil {
		return fmt.Errorf("exec batch: %w", err)
	}
	return nil
}

type FindUserParams struct {
	UserID gocql.UUID
}

type FindUserResult struct {
	UserID    gocql.UUID
	CreatedAt time.Time
	Email     string
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
	if err := q.Scan(&result.UserID, &result.CreatedAt, &result.Email, &result.Username); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &result, nil
}

type ListUserIDsResult struct {
	UserID gocql.UUID
}

type ListUserIDsQuerier struct {
	query  *gocql.Query
	logger log.Logger
}

func (q *ListUserIDsQuerier) All(ctx context.Context) ([]*ListUserIDsResult, error) {
	var results []*ListUserIDsResult
	var pageState []byte
	for {
		page, err := q.Page(ctx, pageState)
		if err != nil {
			return nil, fmt.Errorf("page: %w", err)
		}
		results = append(results, page.Results()...)
		if len(page.PageState()) == 0 {
			break
		}
		pageState = page.PageState()
	}
	return results, nil
}

type ListUserIDsResultsPage struct {
	results   []*ListUserIDsResult
	pageState []byte
	numRows   int
}

func (page *ListUserIDsResultsPage) Results() []*ListUserIDsResult { return page.results }
func (page *ListUserIDsResultsPage) NumRows() int                  { return page.numRows }
func (page *ListUserIDsResultsPage) PageState() []byte             { return page.pageState }

func (q *ListUserIDsQuerier) Page(ctx context.Context, pageState []byte) (*ListUserIDsResultsPage, error) {
	var results []*ListUserIDsResult
	iter := q.query.WithContext(ctx).PageState(pageState).Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			q.logger.Error("iter.Close() returned with error", "error", err)
		}
	}()
	nextPageState := iter.PageState()
	scanner := iter.Scanner()
	for scanner.Next() {
		var result ListUserIDsResult
		if err := scanner.Scan(&result.UserID); err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		results = append(results, &result)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}
	return &ListUserIDsResultsPage{results: results, pageState: nextPageState, numRows: iter.NumRows()}, nil
}

func (c *Client) ListUserIDs(opts ...gocqlc.QueryOption) *ListUserIDsQuerier {
	session := c.Session()
	q := session.Query("SELECT user_id FROM users;")
	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	return &ListUserIDsQuerier{query: q, logger: c.Logger()}
}

type ListUsersResult struct {
	UserID    gocql.UUID
	CreatedAt time.Time
	Email     string
	Username  string
}

type ListUsersQuerier struct {
	query  *gocql.Query
	logger log.Logger
}

func (q *ListUsersQuerier) All(ctx context.Context) ([]*ListUsersResult, error) {
	var results []*ListUsersResult
	var pageState []byte
	for {
		page, err := q.Page(ctx, pageState)
		if err != nil {
			return nil, fmt.Errorf("page: %w", err)
		}
		results = append(results, page.Results()...)
		if len(page.PageState()) == 0 {
			break
		}
		pageState = page.PageState()
	}
	return results, nil
}

type ListUsersResultsPage struct {
	results   []*ListUsersResult
	pageState []byte
	numRows   int
}

func (page *ListUsersResultsPage) Results() []*ListUsersResult { return page.results }
func (page *ListUsersResultsPage) NumRows() int                { return page.numRows }
func (page *ListUsersResultsPage) PageState() []byte           { return page.pageState }

func (q *ListUsersQuerier) Page(ctx context.Context, pageState []byte) (*ListUsersResultsPage, error) {
	var results []*ListUsersResult
	iter := q.query.WithContext(ctx).PageState(pageState).Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			q.logger.Error("iter.Close() returned with error", "error", err)
		}
	}()
	nextPageState := iter.PageState()
	scanner := iter.Scanner()
	for scanner.Next() {
		var result ListUsersResult
		if err := scanner.Scan(&result.UserID, &result.CreatedAt, &result.Email, &result.Username); err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		results = append(results, &result)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}
	return &ListUsersResultsPage{results: results, pageState: nextPageState, numRows: iter.NumRows()}, nil
}

func (c *Client) ListUsers(opts ...gocqlc.QueryOption) *ListUsersQuerier {
	session := c.Session()
	q := session.Query("SELECT * FROM users;")
	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	return &ListUsersQuerier{query: q, logger: c.Logger()}
}
