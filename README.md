# cqlc

[`cqlc`](https://github.com/razcoen/cqlc) is a tool designed to simplify working with Cassandra databases by generating type-safe code from CQL queries. 
<br> Inspired by [`sqlc`](https://github.com/sqlc-dev/sqlc/).


> [!WARNING]
> This library is in its early stages. There are known bugs, incomplete documentation, and potential breaking changes. Contributions and feedback are welcome.

## Features

- **Type-Safe Code Generation:**
  <br><br> Automatically generates code from your CQL queries, ensuring compile-time type safety.
  <br> Currently only **Golang** is supported.

- **Supported CQL Commands:**
  - [x] `SELECT`
  - [x] `INSERT`
  - [x] `DELETE`
  - [x] `CREATE TABLE`
  - [ ] `ALTER TABLE`
  - [ ] `CREATE TYPE`
  - [ ] `CREATE KEYSPACE`
 
- **Supported Query Annotations:**
  <br><br> When defining your CQL queries, you can use the following annotations to specify the expected behavior of each query:
  - `:one` — Fetch a single row.
  - `:many` — Fetch multiple rows.
  - `:exec` — Execute a query without returning rows.
  - `:batch` — Execute a batch of insert operations.


## Installation

To install the `cqlc`, run:

```bash
go install github.com/razcoen/cqlc/cmd/cqlc@latest
```

## Getting Started

1. **Set up your cassandra schema:**

   Ensure that your Cassandra keyspace and tables are properly defined. For example, you might have a `schema.cql` file with the following content:

   #### schema.cql
   ```cql
    CREATE KEYSPACE IF NOT EXISTS your_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
    
    USE your_keyspace;
    
    CREATE TABLE IF NOT EXISTS users (
        user_id    UUID PRIMARY KEY,
        username   TEXT,
        email      TEXT,
        created_at TIMESTAMP
    );
   ```

2. **Define your cql queries:**

   Create a `queries.cql` file containing your CQL queries. For example:

   #### queries.yaml
   ```cql
    -- name: CreateUser :exec
    INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);
    
    -- name: FindUser :one
    SELECT * FROM users WHERE user_id = ? LIMIT 1;
   ```
   
3. **Create a configuration file:**

   In your project directory, create a `cqlc.yaml` configuration file with the following content:

   #### cqlc.yaml
   ```yaml
    cql:
      - queries: "queries.cql"
        schema: "schema.cql"
        gen:
          go:
            package: "gencql"
            out: "gencql"
   ```

   This configuration specifies that the generated code will be placed in the `gencql` directory under the `gencql` package, using `queries.cql` for the CQL queries and `schema.cql` for the database schema.


4. **Generate go code:**

   Run the `cqlc` tool to generate Go code from your CQL queries:

   ```bash
   cqlc generate
   ```

   This command will generate Go files in the `gencql` directory as specified in your configuration.

5. **Integrate into your go application:**

   Use the generated code in your Go application:

   #### main.go
    ```go
    package main
    
    import (
    	"context"
    	"log"
    	"testcql/gencql"
    	"time"
    
    	"github.com/gocql/gocql"
    )
    
    //go:generate cqlc generate
    
    func main() {
    	cluster := gocql.NewCluster("127.0.0.1")
    	cluster.Keyspace = "your_keyspace"
    	session, err := cluster.CreateSession()
    	if err != nil {
    		log.Fatal(err)
    	}
    	defer session.Close()
    
    	client, err := gencql.NewClient(session, nil)
    	if err != nil {
    		log.Fatal(err)
    	}
    	userID := gocql.TimeUUID()
    	err = client.CreateUser(context.Background(), &gencql.CreateUserParams{
    		UserID:    userID,
    		Username:  "John Doe",
    		Email:     "john.doe@example.com",
    		CreatedAt: time.Now(),
    	})
    	if err != nil {
    		log.Fatal(err)
    	}
    
    	user, err := client.FindUser(context.Background(), &gencql.FindUserParams{UserID: userID})
    	if err != nil {
    		log.Fatal(err)
    	}
    
    	log.Printf("User: %+v\n", user)
    }
    ```

## Contributing

Contributions are welcome! 
<br> Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. 
<br> See the [LICENSE](LICENSE) file for details.
