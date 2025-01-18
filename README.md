
[cqlc](https://github.com/razcoen/cqlc) is a Go library heavily inspired by [sqlc](https://github.com/sqlc-dev/sqlc/), designed to simplify working with Cassandra databases by generating type-safe Go code from CQL queries. 

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
