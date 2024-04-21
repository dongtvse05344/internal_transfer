# internal_transfer
Exercise: Internal transfers System with an HTTP Interface


# SQL create table
![img.png](img.png)

```
CREATE TABLE accounts (
  account_id INT PRIMARY KEY,  -- Unique identifier for the user account
  balance DECIMAL(10,2) NOT NULL DEFAULT 0.00  -- Account balance with decimal precision
);

CREATE TABLE transactions (
  transaction_id INT PRIMARY KEY AUTO_INCREMENT,  -- Unique identifier for the transaction (auto-increment)
  from_account_id INT NOT NULL,  -- Account ID initiating the transaction (foreign key to accounts.account_id)
  to_account_id INT NOT NULL,  -- Account ID receiving the transaction (foreign key to accounts.account_id)
  amount DECIMAL(10,2) NOT NULL,  -- Transaction amount with decimal precision
  FOREIGN KEY (from_account_id) REFERENCES accounts(account_id),  -- Enforces foreign key relationship between tables
  FOREIGN KEY (to_account_id) REFERENCES accounts(account_id)  -- Enforces foreign key relationship between tables
);
```