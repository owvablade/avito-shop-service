package txmanager

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type TxManager interface {
	ReadOnly(context.Context, func(context.Context) error) error
	ReadWrite(context.Context, func(context.Context) error) error
}

type SqlxTxManager struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *SqlxTxManager {
	return &SqlxTxManager{db: db}
}

func (tm *SqlxTxManager) ReadOnly(ctx context.Context, fn func(context.Context) error) error {
	txOptions := &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}

	return tm.withTx(ctx, txOptions, fn)
}

func (tm *SqlxTxManager) ReadWrite(ctx context.Context, fn func(context.Context) error) error {
	txOptions := &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false}

	return tm.withTx(ctx, txOptions, fn)
}

func (tm *SqlxTxManager) withTx(ctx context.Context, txOpts *sql.TxOptions, fn func(context.Context) error) error {
	tx, err := tm.db.BeginTxx(ctx, txOpts)
	if err != nil {
		return err
	}

	err = fn(ctx)
	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
