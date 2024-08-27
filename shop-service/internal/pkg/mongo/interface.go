package mongo

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Tx interface {
	pgx.Tx
}

type RepoTx interface {
	Begin(ctx context.Context) (Tx, error)
	TxRollback(ctx context.Context, tx Tx, err error) error
}
