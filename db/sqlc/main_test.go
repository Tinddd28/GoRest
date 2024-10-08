package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var connPool *pgxpool.Pool

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:12345678@localhost:5432/simple_app?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	connPool, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(connPool)
	os.Exit(m.Run())

}
