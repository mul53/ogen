CREATE TABLE `blocks` (
    `block_hash` varchar(255) PRIMARY KEY NOT NULL,
    `block_signature` varchar(255) NOT NULL,
    `block_randao_signature` varchar(255) NOT NULL,
    `height` int NOT NULL
);

CREATE TABLE `block_headers` (
    `block_hash` varchar(255) PRIMARY KEY NOT NULL,
    `version` int NOT NULL,
    `nonce` bigint NOT NULL,
    `tx_merkle_root` varchar(255) NOT NULL,
    `tx_multi_merkle_root` varchar(255) NOT NULL,
    `vote_merkle_root` varchar(255) NOT NULL,
    `deposit_merkle_root` varchar(255) NOT NULL,
    `exit_merkle_root` varchar(255) NOT NULL,
    `vote_slashing_merkle_root` varchar(255) NOT NULL,
    `randao_slashing_merkle_root` varchar(255) NOT NULL,
    `proposer_slashing_merkle_root` varchar(255) NOT NULL,
    `governance_votes_merkle_root` varchar(255) NOT NULL,
    `previous_block_hash` varchar(255) NOT NULL,
    `timestamp` int NOT NULL,
    `slot` int NOT NULL,
    `state_root` varchar(255) NOT NULL,
    `fee_address` varchar(255) NOT NULL
);

CREATE TABLE `votes` (
    `block_hash` varchar(255) NOT NULL,
    `signature` varchar(255) NOT NULL,
    `participation_bitfield` varchar(12518) NOT NULL,
    `data_slot` int NOT NULL,
    `data_from_epoch` int NOT NULL,
    `data_from_hash` varchar(255) NOT NULL,
    `data_to_epoch` int NOT NULL,
    `data_to_hash` varchar(255) NOT NULL,
    `data_beacon_block_hash` varchar(255) NOT NULL,
    `data_nonce` bigint NOT NULL,
    `vote_hash` varchar(255) NOT NULL
);

CREATE TABLE `deposits` (
    `block_hash` varchar(255) NOT NULL,
    `public_key` varchar(255) NOT NULL,
    `signature` varchar(255) NOT NULL,
    `data_public_key` varchar(255) PRIMARY KEY NOT NULL,
    `data_proof_of_possession` varchar(255) NOT NULL,
    `data_withdrawal_address` varchar(255) NOT NULL
);

CREATE TABLE `accounts` (
    `account` varchar(255) PRIMARY KEY NOT NULL,
    `confirmed` bigint DEFAULT 0,
    `unconfirmed` bigint DEFAULT 0,
    `locked` bigint DEFAULT 0,
    `total_sent` bigint DEFAULT 0,
    `total_received` bigint DEFAULT 0
);

CREATE TABLE `validators` (
    `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `public_key` varchar(255) NOT NULL,
    `status` int DEFAULT 0,
    `exit` boolean DEFAULT false,
    `penalized` boolean DEFAULT false,
    `balance` bigint DEFAULT 0,
    `payee_address` varchar(255) NOT NULL,
    `first_active_epoch` int DEFAULT 0,
    `last_active_epoch` int DEFAULT 0
);

CREATE TABLE `exits` (
    `block_hash` varchar(255) NOT NULL,
    `validator_public_key` varchar(255) NOT NULL,
    `withdrawal_public_key` varchar(255) NOT NULL,
    `signature` varchar(255) NOT NULL
);

CREATE TABLE `tx_single` (
    `hash` varchar(255) NOT NULL,
    `block_hash` varchar(255) NOT NULL,
    `tx_type` int NOT NULL,
    `to_addr` varchar(255) NOT NULL,
    `from_public_key` varchar(255) NOT NULL,
    `from_public_key_hash` varchar(255) NOT NULL,
    `amount` bigint NOT NULL,
    `nonce` int NOT NULL,
    `fee` bigint NOT NULL,
    `signature` varchar(255) NOT NULL
);

CREATE TABLE `vote_slashing` (
    `block_hash` varchar(255) NOT NULL,
    `vote_1` varchar(255) NOT NULL,
    `vote_2` varchar(255) NOT NULL
);

CREATE TABLE `randao_slashing` (
    `block_hash` varchar(255) PRIMARY KEY NOT NULL,
    `randao_reveal` varchar(255) NOT NULL,
    `slot` int NOT NULL,
    `validator_public_key` varchar(255) NOT NULL
);

CREATE TABLE `proposer_slashing` (
    `block_hash` varchar(255) PRIMARY KEY NOT NULL,
    `blockheader_1` varchar(255) NOT NULL,
    `blockheader_2` varchar(255) NOT NULL,
    `signature_1` varchar(255) NOT NULL,
    `signature_2` varchar(255) NOT NULL,
    `validator_public_key` varchar(255) NOT NULL
);

CREATE TABLE `slots` (
    `slot` int PRIMARY KEY NOT NULL,
    `block_hash` varchar(255) NOT NULL,
    `proposer_index` int NOT NULL,
    `proposed` boolean NOT NULL
);

CREATE TABLE `epochs` (
    `epoch` int PRIMARY KEY NOT NULL,
    `slot_1` int NOT NULL,
    `slot_2` int NOT NULL,
    `slot_3` int NOT NULL,
    `slot_4` int NOT NULL,
    `slot_5` int NOT NULL,
    `participation_percentage` int NOT NULL,
    `finalized` bool NOT NULL,
    `justified` bool NOT NULL,
    `randao` varchar(255)
);

ALTER TABLE `block_headers` ADD FOREIGN KEY (`block_hash`) REFERENCES `blocks` (`block_hash`);

ALTER TABLE `vote_slashing` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `votes` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `deposits` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `exits` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `proposer_slashing` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `randao_slashing` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `tx_single` ADD FOREIGN KEY (`block_hash`) REFERENCES `block_headers` (`block_hash`);

ALTER TABLE `epochs` ADD FOREIGN KEY (`slot_1`) REFERENCES `slots` (`slot`);

ALTER TABLE `epochs` ADD FOREIGN KEY (`slot_2`) REFERENCES `slots` (`slot`);

ALTER TABLE `epochs` ADD FOREIGN KEY (`slot_3`) REFERENCES `slots` (`slot`);

ALTER TABLE `epochs` ADD FOREIGN KEY (`slot_4`) REFERENCES `slots` (`slot`);

ALTER TABLE `epochs` ADD FOREIGN KEY (`slot_5`) REFERENCES `slots` (`slot`);

ALTER TABLE `validators` ADD FOREIGN KEY (`public_key`) REFERENCES `deposits` (`data_public_key`);

ALTER TABLE `proposer_slashing` ADD FOREIGN KEY (`validator_public_key`) REFERENCES `deposits` (`data_public_key`);

ALTER TABLE `exits` ADD FOREIGN KEY (`validator_public_key`) REFERENCES `deposits` (`data_public_key`);

ALTER TABLE `tx_single` ADD FOREIGN KEY (`from_public_key_hash`) REFERENCES `accounts` (`account`);

CREATE UNIQUE INDEX `blocks_index_0` ON `blocks` (`block_hash`);

CREATE UNIQUE INDEX `blocks_index_1` ON `blocks` (`height`);

CREATE UNIQUE INDEX `block_headers_index_2` ON `block_headers` (`block_hash`);

CREATE UNIQUE INDEX `block_headers_index_3` ON `block_headers` (`slot`);

CREATE UNIQUE INDEX `block_headers_index_4` ON `block_headers` (`timestamp`);

CREATE UNIQUE INDEX `deposits_index_5` ON `deposits` (`data_public_key`);

CREATE UNIQUE INDEX `accounts_index_6` ON `accounts` (`account`);

CREATE UNIQUE INDEX `validators_index_7` ON `validators` (`id`);

CREATE UNIQUE INDEX `exits_index_8` ON `exits` (`validator_public_key`);

CREATE UNIQUE INDEX `tx_single_index_9` ON `tx_single` (`hash`);

CREATE UNIQUE INDEX `slots_index_10` ON `slots` (`slot`);

CREATE UNIQUE INDEX `epochs_index_11` ON `epochs` (`epoch`);
