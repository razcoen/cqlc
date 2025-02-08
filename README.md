# cqlc

![GitHub CI](https://github.com/razcoen/cqlc/actions/workflows/go.yaml/badge.svg) [![codecov](https://codecov.io/gh/razcoen/cqlc/graph/badge.svg?token=RCKM4XXK1I)](https://codecov.io/gh/razcoen/cqlc)

---

[`cqlc`](https://github.com/razcoen/cqlc) is a tool designed to simplify working with Cassandra databases by generating type-safe code from CQL queries.

## Overview

- [Installation](https://razcoen.github.io/cqlc/usage/installation)
- [Getting started](https://razcoen.github.io/cqlc/usage/getting-started)

## Features

- **Type-Safe Code Generation:**
  <br><br> Automatically generates code from your CQL queries, ensuring compile-time type safety.
  <br> Currently only **Golang** is supported.

- **Supported CQL Commands:**
  - [x] `SELECT`
  - [x] `INSERT`
  - [x] `DELETE`
  - [x] `CREATE TABLE`
  - [x] `ALTER TABLE`
    - [x] ADD COLUMN
    - [x] DROP COLUMN
    - [ ] ALTER TYPE
  - [x] `DROP TABLE`
  - [ ] `CREATE TYPE`
  - [ ] `CREATE KEYSPACE`

- **Supported Query Annotations:**
  <br><br> When defining your CQL queries, you can use the following annotations to specify the expected behavior of each query:
  - `:one` — Fetch a single row.
  - `:many` — Fetch multiple rows.
  - `:exec` — Execute a query without returning rows.
  - `:batch` — Execute a batch of insert operations.
