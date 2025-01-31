// Code generated by cqlc. DO NOT EDIT.

package example

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
	CreatedAt time.Time
}

func (c *Client) CreateUser(ctx context.Context, params *CreateUserParams, opts ...gocqlc.QueryOption) error {
	q := c.session.Query("INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);", params.UserID, params.Username, params.Email, params.CreatedAt)
	q = q.WithContext(ctx)
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
	b := c.session.NewBatch(gocql.UnloggedBatch)
	for _, v := range params {
		b.Query("INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);", v.UserID, v.Username, v.Email, v.CreatedAt)
	}
	b = b.WithContext(ctx)
	for _, opt := range opts {
		b = opt.Apply(b)
	}
	if err := c.session.ExecuteBatch(b); err != nil {
		return fmt.Errorf("exec batch: %w", err)
	}
	return nil
}

type DeleteUserParams struct {
	UserID gocql.UUID
}

func (c *Client) DeleteUser(ctx context.Context, params *DeleteUserParams, opts ...gocqlc.QueryOption) error {
	q := c.session.Query("DELETE FROM users WHERE user_id = ?;", params.UserID)
	q = q.WithContext(ctx)
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	if err := q.Exec(); err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	return nil
}

type DeleteUsersParams struct {
	UserID gocql.UUID
}

func (c *Client) DeleteUsers(ctx context.Context, params []*DeleteUsersParams, opts ...gocqlc.BatchOption) error {
	b := c.session.NewBatch(gocql.UnloggedBatch)
	for _, v := range params {
		b.Query("DELETE FROM users WHERE user_id = ?;", v.UserID)
	}
	b = b.WithContext(ctx)
	for _, opt := range opts {
		b = opt.Apply(b)
	}
	if err := c.session.ExecuteBatch(b); err != nil {
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
	q := c.session.Query("SELECT * FROM users WHERE user_id = ? LIMIT 1;", params.UserID)
	q = q.WithContext(ctx)
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	var result FindUserResult
	if err := q.Scan(&result.UserID, &result.CreatedAt, &result.Email, &result.Username); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &result, nil
}

type FindUsersParams struct {
	Email string
}

type FindUsersResult struct {
	UserID    gocql.UUID
	CreatedAt time.Time
	Email     string
	Username  string
}

type FindUsersQuerier struct {
	query  *gocql.Query
	logger gocqlc.Logger
}

func (q *FindUsersQuerier) All(ctx context.Context) ([]*FindUsersResult, error) {
	var results []*FindUsersResult
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

type FindUsersResultsPage struct {
	results   []*FindUsersResult
	pageState []byte
	numRows   int
}

func (page *FindUsersResultsPage) Results() []*FindUsersResult { return page.results }
func (page *FindUsersResultsPage) NumRows() int                { return page.numRows }
func (page *FindUsersResultsPage) PageState() []byte           { return page.pageState }

func (q *FindUsersQuerier) Page(ctx context.Context, pageState []byte) (*FindUsersResultsPage, error) {
	var results []*FindUsersResult
	iter := q.query.WithContext(ctx).PageState(pageState).Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			q.logger.Error("iter.Close() returned with error", "error", err)
		}
	}()
	nextPageState := iter.PageState()
	scanner := iter.Scanner()
	for scanner.Next() {
		var result FindUsersResult
		if err := scanner.Scan(&result.UserID, &result.CreatedAt, &result.Email, &result.Username); err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		results = append(results, &result)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}
	return &FindUsersResultsPage{results: results, pageState: nextPageState, numRows: iter.NumRows()}, nil
}

func (c *Client) FindUsers(params *FindUsersParams, opts ...gocqlc.QueryOption) *FindUsersQuerier {
	q := c.session.Query("SELECT * FROM users WHERE email = ? ALLOW FILTERING;", params.Email)
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	return &FindUsersQuerier{query: q, logger: c.logger}
}
