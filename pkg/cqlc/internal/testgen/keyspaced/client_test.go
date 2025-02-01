package keyspaced

import (
	"context"
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/razcoen/cqlc/internal/testcassandra"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	t.Run("create 2 users and find one by one", func(t *testing.T) {
		client := createClientWithEmptyAuthKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &FindUserParams{UserID: userID1})
		require.NoError(t, err)
		require.Equal(t, FindUserResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *result1)

		result2, err := client.FindUser(ctx, &FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, FindUserResult{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		}, *result2)
	})

	t.Run("create 2 users and find one of many", func(t *testing.T) {
		client := createClientWithEmptyAuthKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		querier := client.FindUsers(&FindUsersParams{Email: "test_email_1"})
		results, err := querier.All(ctx)
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.Equal(t, FindUsersResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *results[0])
	})

	t.Run("create 2 users and find both", func(t *testing.T) {
		client := createClientWithEmptyAuthKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		querier := client.FindUsers(&FindUsersParams{Email: "test_email_1"})
		results, err := querier.All(ctx)
		require.NoError(t, err)
		require.Len(t, results, 2)
		require.ElementsMatch(t, []*FindUsersResult{
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
		client := createClientWithEmptyAuthKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUsers(ctx, []*CreateUsersParams{
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

		result1, err := client.FindUser(ctx, &FindUserParams{UserID: userID1})
		require.NoError(t, err)
		require.Equal(t, FindUserResult{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		}, *result1)
		result2, err := client.FindUser(ctx, &FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, FindUserResult{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		}, *result2)
	})

	t.Run("create 2 users and delete 1 and find the other", func(t *testing.T) {
		client := createClientWithEmptyAuthKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		err = client.DeleteUser(ctx, &DeleteUserParams{UserID: userID1})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &FindUserParams{UserID: userID1})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result1)

		result2, err := client.FindUser(ctx, &FindUserParams{UserID: userID2})
		require.NoError(t, err)
		require.Equal(t, FindUserResult{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		}, *result2)
	})

	t.Run("create 2 users and batch delete both", func(t *testing.T) {
		client := createClientWithEmptyAuthKeyspace(t)

		userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
		createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
		err := client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID1,
			Username:  "test_user_1",
			Email:     "test_email_1",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)
		err = client.CreateUser(ctx, &CreateUserParams{
			UserID:    userID2,
			Username:  "test_user_2",
			Email:     "test_email_2",
			CreatedAt: createdAt,
		})
		require.NoError(t, err)

		err = client.DeleteUsers(ctx, []*DeleteUsersParams{{UserID: userID1}, {UserID: userID2}})
		require.NoError(t, err)

		result1, err := client.FindUser(ctx, &FindUserParams{UserID: userID1})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result1)
		result2, err := client.FindUser(ctx, &FindUserParams{UserID: userID2})
		require.ErrorIs(t, err, gocql.ErrNotFound)
		require.Nil(t, result2)
	})
}

func createClientWithEmptyAuthKeyspace(t *testing.T) *Client {
	session, err := testcassandra.Connect("system")
	require.NoError(t, err, "create cassandra session in system keyspace")

	// Recreate the auth keyspace.
	// There is an implicit dependency on the auth keyspace already existing within the schema file.
	require.NoError(t, session.Session.Query("DROP KEYSPACE IF EXISTS auth;").Exec(), "drop auth keyspace")
	require.NoError(t, session.Session.Query("CREATE KEYSPACE IF NOT EXISTS auth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1};").Exec(), "create auth keyspace")
	require.NoError(t, testcassandra.ExecFile(session.Session, "../../testdata/keyspaced_schema.cql"), "migrate cassandra schema")

	client, err := NewClient(session.Session, nil)
	require.NoError(t, err, "create client")
	t.Cleanup(func() {
		// Make sure to drop the keyspace after the test is done.
		require.NoError(t, session.Session.Query("DROP KEYSPACE IF EXISTS auth;").Exec(), "drop auth keyspace")
		require.NoError(t, session.Close(), "close session")
		// TODO: It might be just better of not encapsulating the close functionality and let the user handle it.
		require.NoError(t, client.Close(), "close client")
	})
	return client
}
