package gocqlc

import (
	"context"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/internal/testcassandra"
	"github.com/razcoen/cqlc/pkg/log"
	"github.com/stretchr/testify/require"
)

func TestQuerier(t *testing.T) {
	sw, err := testcassandra.ConnectWithRandomKeyspace()
	require.NoError(t, err)
	session := sw.Session

	require.NoError(t, session.Query("CREATE TABLE users (user_id uuid PRIMARY KEY, email text, username text, created_at timestamp);").Exec())
	for i := 0; i < 100; i++ {
		userID := gocql.TimeUUID()
		email := fmt.Sprintf("user%d@example.com", i%10)
		username := fmt.Sprintf("user%d", i)
		require.NoError(t, session.Query("INSERT INTO users (user_id, email, username, created_at) VALUES (?, ?, ?, ?);", userID, email, username, time.Now()).Exec())
	}

	type UserRow struct {
		UserID    gocql.UUID
		CreatedAt time.Time
		Email     string
		Username  string
	}
	q := session.Query("SELECT * FROM users WHERE email = ? ALLOW FILTERING;", "user1@example.com")
	scan := func(it *gocql.Iter, dest *UserRow) bool {
		return it.Scan(&(*dest).UserID, &(*dest).CreatedAt, &(*dest).Email, &(*dest).Username)
	}

	t.Run("iter", func(t *testing.T) {
		iter := NewQuerier(q, scan, log.NopLogger()).Iter(context.Background())
		rows := slices.Collect(iter.Rows())
		require.NoError(t, iter.Err())
		require.Len(t, rows, 10)
		usernames := make([]string, len(rows))
		for i, row := range rows {
			usernames[i] = row.Username
		}
		require.ElementsMatch(t, []string{"user1", "user11", "user21", "user31", "user41", "user51", "user61", "user71", "user81", "user91"}, usernames)
	})

	t.Run("iter with small page size", func(t *testing.T) {
		iter := NewQuerier(q, scan, log.NopLogger()).Iter(context.Background(), WithPageSize(2))
		rows := slices.Collect(iter.Rows())
		require.NoError(t, iter.Err())
		require.Len(t, rows, 10)
		usernames := make([]string, len(rows))
		for i, row := range rows {
			usernames[i] = row.Username
		}
		require.ElementsMatch(t, []string{"user1", "user11", "user21", "user31", "user41", "user51", "user61", "user71", "user81", "user91"}, usernames)
	})

	t.Run("page iter without page state", func(t *testing.T) {
		iter := NewQuerier(q, scan, log.NopLogger()).PageIter(context.Background(), nil)
		rows := slices.Collect(iter.Rows())
		require.NoError(t, iter.Err())
		require.Len(t, rows, 10)
		usernames := make([]string, len(rows))
		for i, row := range rows {
			usernames[i] = row.Username
		}
		require.ElementsMatch(t, []string{"user1", "user11", "user21", "user31", "user41", "user51", "user61", "user71", "user81", "user91"}, usernames)
	})

	t.Run("page iter without page state with small page size", func(t *testing.T) {
		iter := NewQuerier(q, scan, log.NopLogger()).PageIter(context.Background(), nil, WithPageSize(2))
		rows := slices.Collect(iter.Rows())
		require.NoError(t, iter.Err())
		require.Len(t, rows, 2)
	})

	t.Run("page iter with small page size and two pages", func(t *testing.T) {
		querier := NewQuerier(q, scan, log.NopLogger())
		iter1 := querier.PageIter(context.Background(), nil, WithPageSize(2))
		rows := slices.Collect(iter1.Rows())
		require.NoError(t, iter1.Err())
		iter2 := querier.PageIter(context.Background(), iter1.Info().PageState(), WithPageSize(2))
		rows = append(rows, slices.Collect(iter2.Rows())...)
		require.NoError(t, iter2.Err())
		require.Len(t, rows, 4)
	})

	t.Run("page iter with small page size loop over all pages", func(t *testing.T) {
		querier := NewQuerier(q, scan, log.NopLogger())
		pageFetches := 0
		var rows []*UserRow
		var pageState []byte
		for {
			iter := querier.PageIter(context.Background(), pageState, WithPageSize(2))
			pageFetches++
			rows = append(rows, slices.Collect(iter.Rows())...)
			require.NoError(t, iter.Err())
			pageState = iter.Info().PageState()
			if pageState == nil {
				break
			}
		}
		require.Len(t, rows, 10)
		require.Equal(t, 6, pageFetches)
		usernames := make([]string, len(rows))
		for i, row := range rows {
			usernames[i] = row.Username
		}
		require.ElementsMatch(t, []string{"user1", "user11", "user21", "user31", "user41", "user51", "user61", "user71", "user81", "user91"}, usernames)
	})

}
