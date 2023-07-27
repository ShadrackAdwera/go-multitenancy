package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var txTestStore TxStore

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	testDbUrl := os.Getenv("TEST_DB_URL")

	pool, err := pgxpool.New(context.Background(), testDbUrl)

	if err != nil {
		log.Fatal("Error establishing connection")
	}

	txTestStore = NewStore(pool)

	defer pool.Close()
	os.Exit(m.Run())
}
