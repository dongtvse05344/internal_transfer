# internal_transfer
Exercise: Internal transfers System with an HTTP Interface
## Quick Start

## Requirement
- Internal transfers application that facilitates financial
  transactions between accounts.
```azure
Consider the currency is the same for all accounts.
* Consider security is not an issue, no need to implement authn or authz
```
- Functional Specifications:
  - HTTP Interface:
    - Create an account
    - Get an account
    - Transfer money between accounts

## Design
### Overview Design
![img_1.png](overview deisgn.png)
### SQL create table
![img.png](img.png)
#### SQL create table
```
CREATE TABLE `accounts` (
                            `id` bigint PRIMARY KEY,
                            `balance` double NOT NULL COMMENT 'must be positive',
                            `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `entries` (
                           `id` bigint PRIMARY KEY AUTO_INCREMENT,
                           `account_id` bigint,
                           `amount` double NOT NULL COMMENT 'can be negative and positive',
                           `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `transfers` (
                             `id` bigint PRIMARY KEY AUTO_INCREMENT,
                             `from_account_id` bigint,
                             `to_account_id` bigint,
                             `amount` double NOT NULL,
                             `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX `accounts_index_0` ON `accounts` (`id`);

CREATE INDEX `entries_index_0` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_1` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`from_account_id`, `to_account_id`);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);
```

### API Design
#### Create an account
```
message CreateAccountRequest {
  int64 id = 1;
  double balance = 2;
}

message CreateAccountResponse {
  int64 code = 1;
  string message = 2;
}
```

#### Get an account
```
message GetAccountRequest {
  int64 id = 1; // from params
}

message GetAccountResponse {
  int64 code = 1;
  string message = 2;
  GetAccountResponseData data = 3;
}

message GetAccountResponseData {
  Account account = 1;
}

message Account {
  int64 id = 1;
  double balance = 2;
}
```

#### Transfer money between accounts
```
message TransferRequest {
    int64 from_account_id = 1;
    int64 to_account_id = 2;
    double amount = 3;
}

message TransferResponse {
    int64 code = 1;
    string message = 2;
}
```

### API
```

service InternalTransfer{
  rpc CreateAccount(CreateAccountRequest) returns (CreateAccountResponse) {
    option (google.api.http) = {
      post: "/v1/account"
      body: "*"
    };
  }

  rpc GetAccount(GetAccountRequest) returns (GetAccountResponse) {
    option (google.api.http) = {
      get: "/v1/account/{id}"
    };
  }

  rpc Transfer(TransferRequest) returns (TransferResponse) {
    option (google.api.http) = {
      post: "/v1/transfer"
      body: "*"
    };
  }
}

```