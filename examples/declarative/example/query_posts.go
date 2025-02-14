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

type ListUserPostsParams struct {
	UserID    gocql.UUID
	CreatedAt time.Time
}

type ListUserPostsResult struct {
	UserID    gocql.UUID
	CreatedAt time.Time
	Content   string
	PostID    gocql.UUID
}

type ListUserPostsQuerier struct {
	query  *gocql.Query
	logger log.Logger
}

func (q *ListUserPostsQuerier) All(ctx context.Context) ([]*ListUserPostsResult, error) {
	var results []*ListUserPostsResult
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

type ListUserPostsResultsPage struct {
	results   []*ListUserPostsResult
	pageState []byte
	numRows   int
}

func (page *ListUserPostsResultsPage) Results() []*ListUserPostsResult { return page.results }
func (page *ListUserPostsResultsPage) NumRows() int                    { return page.numRows }
func (page *ListUserPostsResultsPage) PageState() []byte               { return page.pageState }

func (q *ListUserPostsQuerier) Page(ctx context.Context, pageState []byte) (*ListUserPostsResultsPage, error) {
	var results []*ListUserPostsResult
	iter := q.query.WithContext(ctx).PageState(pageState).Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			q.logger.Error("iter.Close() returned with error", "error", err)
		}
	}()
	nextPageState := iter.PageState()
	scanner := iter.Scanner()
	for scanner.Next() {
		var result ListUserPostsResult
		if err := scanner.Scan(&result.UserID, &result.CreatedAt, &result.Content, &result.PostID); err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		results = append(results, &result)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}
	return &ListUserPostsResultsPage{results: results, pageState: nextPageState, numRows: iter.NumRows()}, nil
}

func (c *Client) ListUserPosts(params *ListUserPostsParams, opts ...gocqlc.QueryOption) *ListUserPostsQuerier {
	session := c.Session()
	q := session.Query("SELECT * FROM posts WHERE user_id = ? AND created_at = ?;", params.UserID, params.CreatedAt)
	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	return &ListUserPostsQuerier{query: q, logger: c.Logger()}
}
