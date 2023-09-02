package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxStore interface {
	Querier
	CreateTenantUserTx(ctx context.Context, input TenantUserTxInput) (TenantUserTxOutput, error)
}

type Store struct {
	*Queries
	pgxpool            *pgxpool.Pool
	tenantDatabasePool map[string]*pgxpool.Pool // Map tenant ID to the respective database pool
	poolLock           *sync.Mutex
}

func NewStore(pool *pgxpool.Pool, tenantDatabasePool map[string]*pgxpool.Pool, poolLock *sync.Mutex) TxStore {
	return &Store{
		pgxpool:            pool,
		Queries:            New(pool),
		tenantDatabasePool: tenantDatabasePool,
		poolLock:           poolLock,
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
