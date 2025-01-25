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
	session, _ := testcassandra.ConnectWithRandomKeyspace(t)
	testcassandra.Exec(t, session, "schema.cql")

	client, err := example.NewClient(session, nil)
	require.NoError(t, err)
	defer func() { require.NoError(t, client.Close()) }()

	ctx := context.Background()
	userID1 := gocql.UUID(uuid.Must(uuid.NewUUID()))
	createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
	err = client.CreateUser(ctx, &example.CreateUserParams{
		UserID:    userID1,
		Username:  "test_user_1",
		Email:     "test_email_1",
		CreatedAt: createdAt,
	})
	require.NoError(t, err)

	userID2 := gocql.UUID(uuid.Must(uuid.NewUUID()))
	err = client.CreateUser(ctx, &example.CreateUserParams{
		UserID:    userID2,
		Username:  "test_user_2",
		Email:     "test_email_2",
		CreatedAt: createdAt,
	})
	require.NoError(t, err)

	result, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID1})
	require.NoError(t, err)
	require.Equal(t, example.FindUserResult{
		UserID:    userID1,
		Username:  "test_user_1",
		Email:     "test_email_1",
		CreatedAt: createdAt,
	}, *result)
}
