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
	// TODO: Requires cassandra cluster running within the CI
	t.Skip()
	session, _ := testcassandra.NewRandomKeyspaceSession(t)
	testcassandra.Migrate(t, session, "schema.cql")
	client, err := example.NewClient(session, nil)
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	uuid, err := uuid.NewUUID()
	require.NoError(t, err)
	userID := gocql.UUID(uuid)
	require.NoError(t, err)
	createdAt := time.Now()
	err = client.CreateUser(ctx, &example.CreateUserParams{
		UserID:    userID,
		Username:  "test_user",
		Email:     "test_email",
		CreatedAt: createdAt,
	})
	require.NoError(t, err)
	result, err := client.FindUsernameByUserID(ctx, &example.FindUsernameByUserIDParams{UserID: userID})
	require.NoError(t, err)
	require.Equal(t, "test_user", result.Username)
	// TODO: Not working since "*" does not match the column ordering.
	//result, err := client.FindUser(ctx, &example.FindUserParams{UserID: userID})
	//require.NoError(t, err)
	//require.Equal(t, example.FindUserResult{
	//	UserID:    userID,
	//	Username:  "test_user",
	//	Email:     "test_email",
	//	CreatedAt: createdAt,
	//}, *result)
}
