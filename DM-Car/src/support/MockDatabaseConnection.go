package support

import (
	"context"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
)

// A common interface for the real and mocked database connection
type PGXInterface interface {
	Ping(ctx context.Context) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close(ctx context.Context) error
}

func CreateMockDatabaseConnection(t *testing.T) pgxmock.PgxConnIface {
	mockDatabaseConnection, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Contrary to the real database connection, the mocked database connection
	// is responsible for closing itself
	defer mockDatabaseConnection.Close(context.Background())
	return mockDatabaseConnection
}

// Helper function that validates that all expected SQL statements have been executed
func ExpectExpectationsToBeMet(mockDatabaseConnection pgxmock.PgxConnIface, t *testing.T) {
	if err := mockDatabaseConnection.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
