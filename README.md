
[cqlc](https://github.com/razcoen/cqlc) designed to simplify working with Cassandra databases by generating type-safe code from CQL queries. Heavily inspired by [sqlc](https://github.com/sqlc-dev/sqlc/).

> [!WARNING]
> The library is in **very early** stages.
> 
> There are many known bugs, missing documentation and there **will be breaking changes**.

## Installation

```bash
go install github.com/razcoen/cqlc/cmd/cqlc@latest
```

## Support Matrix

### Cassandra Syntax

* [x] SELECT
* [x] INSERT
* [x] DELETE
* [x] CREATE TABLE
* [ ] ALTER TABLE
* [ ] CREATE TYPE
* [ ] CREATE KEYSPACE
* [ ] CREATE MATERIALIZED VIEW

### Golang

* [x] `:one` Fetch one row
* [x] `:many` Fetch many rows
* [x] `:exec` Execute query
* [x] `:batch` Batch insert 
