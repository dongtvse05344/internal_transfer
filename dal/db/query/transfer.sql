-- name: CreateTransfer :execresult
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
             ?,?,?
         );

-- name: GetTransferById :one
SELECT id,
       from_account_id,
       to_account_id,
       amount
  FROM transfers
 WHERE id =  ?;
