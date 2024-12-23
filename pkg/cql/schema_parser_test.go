package cql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestSchemaParser(t *testing.T) {
	// List of test cases for schema-related queries
	tests := []struct {
		query          string
		expectedErr    bool
		expectedSchema *Schema
	}{
		// Valid CREATE TABLE
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT, age INT);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: &DataType{NativeType: ptr(NativeTypeUUID)}},
						{Name: "name", DataType: &DataType{NativeType: ptr(NativeTypeText)}},
						{Name: "age", DataType: &DataType{NativeType: ptr(NativeTypeInt)}},
					}},
				}},
			}},
		},
		// Valid CREATE TABLE with compound primary key
		{
			query:       "CREATE TABLE orders (order_id UUID, customer_id UUID, total DECIMAL, PRIMARY KEY (order_id, customer_id));",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "orders", Columns: []*Column{
						{Name: "order_id", DataType: &DataType{NativeType: ptr(NativeTypeUUID)}},
						{Name: "customer_id", DataType: &DataType{NativeType: ptr(NativeTypeUUID)}},
						{Name: "total", DataType: &DataType{NativeType: ptr(NativeTypeDecimal)}},
					}},
				}},
			}},
		},
		// Valid CREATE TABLE with a collection (SET)
		{
			query:       "CREATE TABLE blog_posts (id UUID PRIMARY KEY, title TEXT, tags SET<TEXT>);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "blog_posts", Columns: []*Column{
						{Name: "id", DataType: &DataType{NativeType: ptr(NativeTypeUUID)}},
						{Name: "title", DataType: &DataType{NativeType: ptr(NativeTypeText)}},
						{Name: "tags", DataType: &DataType{CollectionType: &CollectionType{Set: &CollectionTypeSet{T: NativeTypeText}}}},
					}},
				}},
			}},
		},
		// Valid CREATE TABLE with a collection (LIST)
		{
			query:       "CREATE TABLE events (id UUID PRIMARY KEY, event_dates LIST<TIMESTAMP>);",
			expectedErr: false,
		},
		// CREATE TABLE with frozen collection (frozen SET)
		{
			query:       "CREATE TABLE events (id UUID PRIMARY KEY, participants FROZEN<SET<TEXT>>);",
			expectedErr: false,
		},
		// CREATE TABLE with frozen collection (frozen LIST)
		{
			query:       "CREATE TABLE orders (id UUID PRIMARY KEY, products FROZEN<LIST<TEXT>>);",
			expectedErr: false,
		},
		// CREATE TABLE with custom data types (UDT)
		{
			query:       "CREATE TYPE address (street TEXT, city TEXT, zip_code INT);",
			expectedErr: false,
		},
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT, address FROZEN<address>);",
			expectedErr: false,
		},
		// CREATE TABLE with clustering keys
		{
			query:       "CREATE TABLE sensor_data (sensor_id UUID, timestamp TIMESTAMP, temperature DECIMAL, PRIMARY KEY (sensor_id, timestamp));",
			expectedErr: false,
		},
		// CREATE TABLE with multi-column clustering key (compound clustering key)
		{
			query:       "CREATE TABLE articles (author TEXT, category TEXT, published_at TIMESTAMP, title TEXT, PRIMARY KEY (author, category, published_at));",
			expectedErr: false,
		},
		// CREATE TABLE with time-to-live (TTL)
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT) WITH TTL 86400;",
			expectedErr: false,
		},
		// CREATE TABLE with additional options (e.g., compaction settings)
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT) WITH compaction = {'class': 'LeveledCompactionStrategy'};",
			expectedErr: false,
		},
		// CREATE MATERIALIZED VIEW
		{
			query:       "CREATE MATERIALIZED VIEW user_by_name AS SELECT id, name FROM users WHERE name IS NOT NULL PRIMARY KEY (name, id);",
			expectedErr: false,
		},
		// CREATE INDEX on a column
		{
			query:       "CREATE INDEX ON users(name);",
			expectedErr: false,
		},
		// DROP TABLE
		{
			query:       "DROP TABLE users;",
			expectedErr: false,
		},
		// Invalid CQL - Missing Column Definition in CREATE TABLE
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, );",
			expectedErr: false,
		},
		// Invalid CQL - Missing Parenthesis
		{
			query:       "CREATE TABLE users id UUID PRIMARY KEY;",
			expectedErr: false,
		},
		// Invalid CREATE TABLE with an invalid data type
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name INVALIDTYPE);",
			expectedErr: false,
		},
		// DROP INDEX
		{
			query:       "DROP INDEX idx_name;",
			expectedErr: false,
		},
		// Invalid DROP TABLE (non-existent table)
		{
			query:       "DROP TABLE nonexistent_table;",
			expectedErr: false,
		},
		// CREATE TABLE with a primary key that uses multiple columns (composite key)
		{
			query:       "CREATE TABLE customer_orders (customer_id UUID, order_id UUID, product TEXT, PRIMARY KEY (customer_id, order_id));",
			expectedErr: false,
		},
		// CREATE TABLE with a complex column type (frozen set of tuples)
		{
			query:       "CREATE TABLE orders (id UUID PRIMARY KEY, items FROZEN<SET<TUPLE<UUID, TEXT>>>);",
			expectedErr: false,
		},
		// CREATE TABLE with frozen set of tuples and TTL
		{
			query:       "CREATE TABLE orders (id UUID PRIMARY KEY, items FROZEN<SET<TUPLE<UUID, TEXT>>>) WITH TTL 3600;",
			expectedErr: false,
		},
		// CREATE TABLE with static columns
		{
			query:       "CREATE TABLE sensor_data (sensor_id UUID, timestamp TIMESTAMP, temperature DECIMAL, PRIMARY KEY (sensor_id, timestamp)) WITH STATIC temperature;",
			expectedErr: false,
		},
		// CREATE TABLE with a compound clustering key and order
		{
			query:       "CREATE TABLE articles (author TEXT, category TEXT, published_at TIMESTAMP, title TEXT, PRIMARY KEY (author, category, published_at)) WITH CLUSTERING ORDER BY (published_at DESC);",
			expectedErr: false,
		},
		// CREATE TABLE with default values
		{
			query:       "CREATE TABLE products (id UUID PRIMARY KEY, price DECIMAL DEFAULT 0.0);",
			expectedErr: false,
		},
		// Invalid CREATE MATERIALIZED VIEW due to missing WHERE clause
		{ // TODO: Unsupported
			query:       "CREATE MATERIALIZED VIEW mv_user_by_name AS SELECT id, name FROM users PRIMARY KEY (name, id);",
			expectedErr: false,
		},
		// CREATE TABLE with custom validation
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT) WITH validation = {'validator': 'my_custom_validator'};",
			expectedErr: false,
		},
		{
			query: `
CREATE TABLE users (id UUID PRIMARY KEY, name TEXT, age INT);
CREATE TABLE orders (order_id UUID, customer_id UUID, total DECIMAL, PRIMARY KEY (order_id, customer_id));
CREATE TABLE logins (id UUID PRIMARY KEY, timestamp TIMESTAMP);
      `,
			expectedErr: false,
		},
	}
	parser := &SchemaParser{}
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			schema, err := parser.Parse(tt.query)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Expected error: %v, but got: %v for query: %s", tt.expectedErr, err, tt.query)
			}
			if tt.expectedSchema != nil && !reflect.DeepEqual(tt.expectedSchema, schema) {
        b1, _ := json.Marshal(tt.expectedSchema)
        b2, _ := json.Marshal(schema)
				t.Errorf("Expected schema: %v, but got: %v for query: %s", tt.expectedSchema, schema, tt.query)
        fmt.Println(string(b1))
        fmt.Println(string(b2))
			}
		})
	}
}

func ptr(nt NativeType) *NativeType {
	return &nt
}
