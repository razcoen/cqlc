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

func TestLogic(t *testing.T) {
	session, _ := testcassandra.ConnectWithRandomKeyspace(t)
	testcassandra.Exec(t, session, "schema.cql")

	client, err := example.NewClient(session, nil)
	require.NoError(t, err)
	defer func() { require.NoError(t, client.Close()) }()

	ctx := context.Background()
	userID := gocql.UUID(uuid.Must(uuid.NewUUID()))
	require.NoError(t, err)

	createdAt := time.UnixMilli(time.Now().UnixMilli()).UTC() // Cassandra only keeps milliseconds precision on timestamps
	err = client.CreateUser(ctx, &example.CreateUserParams{
		UserID:    userID,
		Username:  "test_user",
		Email:     "test_email",
		CreatedAt: createdAt,
	})
	require.NoError(t, err)

	result, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID})
	require.NoError(t, err)
	require.Equal(t, example.FindUserResult{
		UserID:    userID,
		Username:  "test_user",
		Email:     "test_email",
		CreatedAt: createdAt,
	}, *result)
}
