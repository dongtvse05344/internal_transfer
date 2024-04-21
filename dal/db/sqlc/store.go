package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (s *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        float64
}

type TransferTxResult struct {
	Transfer    sql.Result
	FromAccount GetAccountRow
	ToAccount   GetAccountRow
	FromEntry   sql.Result
	ToEntry     sql.Result
}

// TransferTx performs a money transfer from one account to the other.
// step 1: query and lock the fromAccount
// step 2: query toAccount
// step 3: create transfer record
// step 4: create transfer entries
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := s.ExecTx(ctx, func(q *Queries) error {
		var err error

		// oder lock by key
		if arg.ToAccountID < arg.FromAccountID {
			_, err = q.GetAccountForUpdate(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}
		}
		_, err = q.GetAccountForBalanceUpdate(ctx, GetAccountForBalanceUpdateParams{
			ID:      arg.FromAccountID,
			Balance: arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: sql.NullInt64{
				Int64: arg.FromAccountID,
				Valid: true,
			},
			ToAccountID: sql.NullInt64{
				Int64: arg.ToAccountID,
				Valid: true,
			},
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}

		_, err = q.UpdateBalance(ctx, UpdateBalanceParams{
			Balance: arg.Amount * -1,
			ID:      arg.FromAccountID,
		})
		if err != nil {
			return err
		}
		_, err = q.UpdateBalance(ctx, UpdateBalanceParams{
			Balance: arg.Amount,
			ID:      arg.ToAccountID,
		})
		if err != nil {
			return err
		}
		result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}
		result.FromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: sql.NullInt64{
				Int64: arg.FromAccountID,
				Valid: true,
			},
			Amount: arg.Amount * -1,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: sql.NullInt64{
				Int64: arg.ToAccountID,
				Valid: true,
			},
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return TransferTxResult{}, err
	}
	return result, nil
}
