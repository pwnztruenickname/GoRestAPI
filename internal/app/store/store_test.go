package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(n *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "host=localhost port=5432 dbname=restapitest user=postgres password=postgres sslmode=disable"
	}

	os.Exit(n.Run())
}
