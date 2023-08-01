package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxStore interface {
	Querier
	CreateTenantUserTx(ctx context.Context, input TenantUserTxInput) (TenantUserTxOutput, error)
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
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.pgxpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rErr := tx.Rollback(ctx); rErr != nil {
			return fmt.Errorf("rollbacc error %w", rErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
