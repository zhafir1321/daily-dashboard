-- Daily Dashboard — MySQL schema
-- Run once:  mysql -u root -p < schema.sql

CREATE DATABASE IF NOT EXISTS daily_dashboard
  CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE daily_dashboard;

CREATE TABLE IF NOT EXISTS todos (
  id        VARCHAR(24) PRIMARY KEY,
  text      VARCHAR(500) NOT NULL,
  due       DATE NULL,
  priority  ENUM('low','medium','high') NOT NULL DEFAULT 'medium',
  done      BOOLEAN NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS events (
  id     VARCHAR(24) PRIMARY KEY,
  date   DATE NOT NULL,
  time   VARCHAR(5) NULL,          -- 'HH:MM' or NULL for all-day
  text   VARCHAR(500) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_events_date (date)
);

CREATE TABLE IF NOT EXISTS transactions (
  id        VARCHAR(24) PRIMARY KEY,
  type      ENUM('income','expense') NOT NULL,
  descr     VARCHAR(500) NOT NULL,
  category  VARCHAR(120) NOT NULL DEFAULT 'Uncategorized',
  amount    DECIMAL(15,2) NOT NULL,
  date      DATE NOT NULL,
  recurring BOOLEAN NOT NULL DEFAULT 0,
  freq      ENUM('daily','weekly','monthly') NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_txn_date (date),
  INDEX idx_txn_type (type)
);
