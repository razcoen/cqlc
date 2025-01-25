package programmatic

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/razcoen/cqlc/examples/programmatic/example"
	"github.com/razcoen/cqlc/pkg/testcassandra"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	ctx := context.Background()

	t.Run("create 2 users and find one by one", func(t *testing.T) {
		client := newClientWithRandomKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID1})
		require.NoError(t, err)
		require.Equal(t, example.FindUserResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *result1)

		result2, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, example.FindUserResult{
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
		err := client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		querier := client.FindUsers(&example.FindUsersParams{Email: "test_email_1"})
		results, err := querier.All(ctx)
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.Equal(t, example.FindUsersResult{
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
		err := client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		querier := client.FindUsers(&example.FindUsersParams{Email: "test_email_1"})
		results, err := querier.All(ctx)
		require.NoError(t, err)
		require.Len(t, results, 2)
		require.ElementsMatch(t, []*example.FindUsersResult{
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
		err := client.CreateUsers(ctx, []*example.CreateUsersParams{
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

		result1, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID1})
		require.NoError(t, err)
		require.Equal(t, example.FindUserResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *result1)
		result2, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, example.FindUserResult{
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
		err := client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		err = client.DeleteUser(ctx, &example.DeleteUserParams{UserID: userID1})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID1})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result1)

		result2, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, example.FindUserResult{
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
		err := client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &example.CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		err = client.DeleteUsers(ctx, []*example.DeleteUsersParams{{UserID: userID1}, {UserID: userID2}})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID1})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result1)
		result2, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID2})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result2)
	})
}

func newClientWithRandomKeyspace(t *testing.T) *example.Client {
	session, _ := testcassandra.ConnectWithRandomKeyspace(t)
	testcassandra.Exec(t, session, "schema.cql")
	client, err := example.NewClient(session, nil)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Close()) })
	return client
}
