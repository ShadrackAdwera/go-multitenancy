package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ShadrackAdwera/go-multitenancy/internal/api"
	db "github.com/ShadrackAdwera/go-multitenancy/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	tenantDatabasePool map[string]*pgxpool.Pool // Map tenant ID to the respective database pool
	poolLock           *sync.Mutex              // Mutex to ensure safe concurrent access to tenantDatabasePool
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("TEST_DB_URL")

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error creating connection pool")
	}

	tenantDatabasePool = make(map[string]*pgxpool.Pool)

	poolLock = &sync.Mutex{}

	store := db.NewStore(pool, tenantDatabasePool, poolLock)

	srv := api.NewServer(store)

	srvAddress := os.Getenv("SERVER_ADDRESS")

	if err := srv.Start(srvAddress); err != nil {
		fmt.Println(err)
		log.Fatal("Error starting server")
	}
}
