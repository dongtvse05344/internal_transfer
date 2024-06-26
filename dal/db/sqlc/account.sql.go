// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO accounts (
    id,
    balance
) VALUES (
             ?, ?
         )
`

type CreateAccountParams struct {
	ID      int64
	Balance float64
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount, arg.ID, arg.Balance)
}

const getAccount = `-- name: GetAccount :one
SELECT id, balance FROM accounts
WHERE id = ? LIMIT 1
`

type GetAccountRow struct {
	ID      int64
	Balance float64
}

func (q *Queries) GetAccount(ctx context.Context, id int64) (GetAccountRow, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i GetAccountRow
	err := row.Scan(&i.ID, &i.Balance)
	return i, err
}

const getAccountForBalanceUpdate = `-- name: GetAccountForBalanceUpdate :one
SELECT id FROM accounts
WHERE id = ? AND balance >= ? LIMIT 1 FOR UPDATE
`

type GetAccountForBalanceUpdateParams struct {
	ID      int64
	Balance float64
}

func (q *Queries) GetAccountForBalanceUpdate(ctx context.Context, arg GetAccountForBalanceUpdateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountForBalanceUpdate, arg.ID, arg.Balance)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id FROM accounts
WHERE id = ?  LIMIT 1 FOR UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
	err := row.Scan(&id)
	return id, err
}

const updateBalance = `-- name: UpdateBalance :execresult
UPDATE accounts
set balance = balance + ?
WHERE id = ?

`

type UpdateBalanceParams struct {
	Balance float64
	ID      int64
}

func (q *Queries) UpdateBalance(ctx context.Context, arg UpdateBalanceParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateBalance, arg.Balance, arg.ID)
}
