
## Support Matrix

### Queries

| Query  | Support | Examples                        |
|--------|------|---------------------------------|
| SELECT | 🟩 Full | `SELECT x FROM y WHERE z = ?;`  |
| INSERT | 🟩 Full | `INSERT INTO x (y) VALUES (z);` |
| DELETE | 🟩 Full | `DELETE FROM x WHERE y = ?;`    |

### Schema

| Query        | Support | Examples                        |
|--------------|---------|---------------------------------|
| CREATE TABLE | 🟩 Full | `CREATE TABLE x (y);`  |
| CREATE KEYSPACE  | 🟥 None  |   |
| CREATE TYPE  | 🟥 None  |   |
| ALTER TABLE  | 🟥 None  |   |
