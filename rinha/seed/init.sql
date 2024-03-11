SET NAMES utf8;
SET time_zone = '-03:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE `bank`;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` tinyint unsigned NOT NULL AUTO_INCREMENT,
  `bal` bigint NOT NULL,
  `lim` int unsigned NOT NULL,
  `ltx` timestamp NOT NULL DEFAULT 0 ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `uuid` varchar(36) DEFAULT (UUID()),
  `acc_id` tinyint unsigned NOT NULL,
  `amt` bigint NOT NULL,
  `typ` char(1) NOT NULL,
  `des` varchar(10) NOT NULL,
  `ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uuid`),
  KEY `acc_id` (`acc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Initial accounts
INSERT INTO `accounts` (`id`, `bal`, `lim`) VALUES
  (1, 0, 100000),
  (2, 0, 80000),
  (3, 0, 1000000),
  (4, 0, 10000000),
  (5, 0, 500000);
