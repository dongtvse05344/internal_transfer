-- name: CreateAccount :execresult
INSERT INTO accounts (
    id,
    balance
) VALUES (
             ?, ?
         );

-- name: GetAccount :one
SELECT id, balance FROM accounts
WHERE id = ? LIMIT 1;

-- name: GetAccountForBalanceUpdate :one
SELECT id FROM accounts
WHERE id = ? AND balance >= ? LIMIT 1 FOR UPDATE;

-- name: GetAccountForUpdate :one
SELECT id FROM accounts
WHERE id = ?  LIMIT 1 FOR UPDATE;

-- name: UpdateBalance :execresult
UPDATE accounts
set balance = balance + ?
WHERE id = ?

--