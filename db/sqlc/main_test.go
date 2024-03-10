package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var queries *Queries

func TestMain(m *testing.M) {
	pool, err := pgxpool.New(context.Background(), "postgresql://user:pass@localhost:5432/sequence_email")
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	queries = New(pool)

	os.Exit(m.Run())
}
