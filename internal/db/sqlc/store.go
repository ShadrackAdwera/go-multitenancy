package db

import "github.com/jackc/pgx/v5/pgxpool"

type TxStore interface {
	Querier
}

type Store struct {
	*Queries
	pgxpool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) TxStore {
	return &Store{
		pgxpool: pool,
		Queries: New(pool),
	}
}

// write execTx
