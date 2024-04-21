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
