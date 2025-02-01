package basic

import (
	"context"
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/razcoen/cqlc/e2e/basic/gen/basic"
	"github.com/razcoen/cqlc/internal/testcassandra"
	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	err := cqlc.Generate(&config.Config{
		CQL: []*config.CQL{
			{
				Queries: "queries.cql",
				Schema:  "schema.cql",
				Gen: &config.CQLGen{
					Overwrite: true,
					Go: &golang.Options{
						Package: "basic",
						Out:     "gen/basic",
					},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	t.Run("create 2 users and find one by one", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID1})
		require.NoError(t, err)
		require.Equal(t, basic.FindUserResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *result1)

		result2, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, basic.FindUserResult{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		}, *result2)
	})

	t.Run("create 2 users and find one of many", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		querier := client.FindUsers(&basic.FindUsersParams{Email: "test_email_1"})
		results, err := querier.All(ctx)
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.Equal(t, basic.FindUsersResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *results[0])
	})

	t.Run("create 2 users and find both", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		querier := client.FindUsers(&basic.FindUsersParams{Email: "test_email_1"})
		results, err := querier.All(ctx)
		require.NoError(t, err)
		require.Len(t, results, 2)
		require.ElementsMatch(t, []*basic.FindUsersResult{
			{
				UserID:    userID1,
				Username:  "test_user_1",
				Email:     "test_email_1",
				CreatedAt: createdAt,
			},
			{
				UserID:    userID2,
				Username:  "test_user_2",
				Email:     "test_email_1",
				CreatedAt: createdAt,
			},
		}, results)
	})

	t.Run("batch create 2 users and find both", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUsers(ctx, []*basic.CreateUsersParams{
			{
				UserID:    userID1,
				Username:  "test_user_1",
				Email:     "test_email_1",
				CreatedAt: createdAt,
			},
			{
				UserID:    userID2,
				Username:  "test_user_2",
				Email:     "test_email_2",
				CreatedAt: createdAt,
			},
		})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID1})
		require.NoError(t, err)
		require.Equal(t, basic.FindUserResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *result1)
		result2, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, basic.FindUserResult{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		}, *result2)
	})

	t.Run("create 2 users and delete 1 and find the other", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		err = client.DeleteUser(ctx, &basic.DeleteUserParams{UserID: userID1})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID1})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result1)

		result2, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, basic.FindUserResult{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		}, *result2)
	})

	t.Run("create 2 users and batch delete both", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &basic.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		err = client.DeleteUsers(ctx, []*basic.DeleteUsersParams{{UserID: userID1}, {UserID: userID2}})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID1})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result1)
		result2, err := client.FindUser(ctx, &basic.FindUserParams{UserID: userID2})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result2)
	})
}

func newClientWithRandomKeyspace(t *testing.T) *basic.Client {
	session, err := testcassandra.ConnectWithRandomKeyspace()
	require.NoError(t, err, "create cassandra session in random keyspace")
	err = testcassandra.ExecFile(session.Session, "schema.cql")
	require.NoError(t, err, "migrate cassandra schema")
	client, err := basic.NewClient(session.Session, nil)
	require.NoError(t, err, "create client")
	t.Cleanup(func() {
		require.NoError(t, session.Close(), "close session")
		// TODO: It might be just better of not encapsulating the close functionality and let the user handle it.
		require.NoError(t, client.Close(), "close client")
	})
	return client
}
