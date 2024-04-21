-- name: CreateEntry :execresult
INSERT INTO entries (
    account_id,
    amount
) VALUES (
             ?,?
         );

