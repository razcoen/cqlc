package cqlc

import (
	"fmt"
	"testing"

	"github.com/razcoen/cqlc/pkg/gocqlhelpers"
	"github.com/stretchr/testify/require"
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
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeText()},
						{Name: "age", DataType: gocqlhelpers.NewTypeInt()},
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
						{Name: "order_id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "customer_id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "total", DataType: gocqlhelpers.NewTypeDecimal()},
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
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "title", DataType: gocqlhelpers.NewTypeText()},
						{Name: "tags", DataType: gocqlhelpers.NewTypeSet(gocqlhelpers.NewTypeText())},
					}},
				}},
			}},
		},
		// Valid CREATE TABLE with a collection (LIST)
		{
			query:       "CREATE TABLE events (id UUID PRIMARY KEY, event_dates LIST<TIMESTAMP>);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "events", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "event_dates", DataType: gocqlhelpers.NewTypeList(gocqlhelpers.NewTypeTimestamp())},
					}},
				}},
			}},
		},
		// CREATE TABLE with frozen collection (frozen SET)
		{
			query:       "CREATE TABLE events (id UUID PRIMARY KEY, participants FROZEN<SET<TEXT>>);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "events", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "participants", DataType: gocqlhelpers.NewTypeSet(gocqlhelpers.NewTypeText())},
					}},
				}},
			}},
		},

		// CREATE TABLE with frozen collection (frozen LIST)
		{
			query:       "CREATE TABLE orders (id UUID PRIMARY KEY, products FROZEN<LIST<TEXT>>);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "orders", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "products", DataType: gocqlhelpers.NewTypeList(gocqlhelpers.NewTypeText())},
					}},
				}},
			}},
		},

		// CREATE TABLE with custom data types (UDT)
		// TODO: Unsupported UDT

		// 		{
		// 			query: `
		// CREATE TYPE address (street TEXT, city TEXT, zip_code INT);
		// CREATE TABLE users (id UUID PRIMARY KEY, name TEXT, address FROZEN<address>);
		//       `,
		// 			expectedErr: false,
		// 			expectedSchema: &Schema{Keyspaces: []*Keyspace{
		// 				{
		// 					Name: defaultKeyspaceName,
		// 					UserDefinedTypes: []*UserDefinedType{
		// 						{Name: "address", Fields: []*UserDefinedTypeField{
		// 							{Name: "street", DataType: NativeTypeText.IntoDataType()},
		// 							{Name: "city", DataType: NativeTypeText.IntoDataType()},
		// 							{Name: "zip_code", DataType: NativeTypeInt.IntoDataType()},
		// 						}},
		// 					},
		// 					Tables: []*Table{
		// 						{Name: "users", Columns: []*Column{
		// 							{Name: "id", DataType: NativeTypeUUID.IntoDataType()},
		// 							{Name: "name", DataType: NativeTypeText.IntoDataType()},
		// 							{Name: "address", DataType: FrozenType{
		// 								DataType: UserDefinedType{Name: "address", Fields: []*UserDefinedTypeField{
		// 									{Name: "street", DataType: NativeTypeText.IntoDataType()},
		// 									{Name: "city", DataType: NativeTypeText.IntoDataType()},
		// 									{Name: "zip_code", DataType: NativeTypeInt.IntoDataType()},
		// 								}}.IntoDataType(),
		// 							}.IntoDataType(),
		// 							}}},
		// 					}},
		// 			}},
		// 		},

		// CREATE TABLE with clustering keys
		// TODO: Unsupported reserved keyword "timestamp"

		// {
		// 	query:       "CREATE TABLE sensor_data (sensor_id UUID, timestamp TIMESTAMP, temperature DECIMAL, PRIMARY KEY (sensor_id, timestamp));",
		// 	expectedErr: false,
		// 	expectedSchema: &Schema{Keyspaces: []*Keyspace{
		// 		{Name: defaultKeyspaceName, Tables: []*Table{
		// 			{Name: "sensor_data", Columns: []*Column{
		// 				{Name: "sensor_id", DataType: NativeTypeUUID.IntoDataType()},
		// 				{Name: "timestamp", DataType: NativeTypeTimestamp.IntoDataType()},
		// 				{Name: "temperature", DataType: NativeTypeDecimal.IntoDataType()},
		// 			}},
		// 		}},
		// 	}},
		// },

		// CREATE TABLE with multi-column clustering key (compound clustering key)
		{
			query:       "CREATE TABLE articles (author TEXT, category TEXT, published_at TIMESTAMP, title TEXT, PRIMARY KEY (author, category, published_at));",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "articles", Columns: []*Column{
						{Name: "author", DataType: gocqlhelpers.NewTypeText()},
						{Name: "category", DataType: gocqlhelpers.NewTypeText()},
						{Name: "published_at", DataType: gocqlhelpers.NewTypeTimestamp()},
						{Name: "title", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
			}},
		},

		// CREATE TABLE with time-to-live (TTL)
		// TODO: Unsupported TTL
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT) WITH TTL 86400;",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
			}},
		},

		// CREATE TABLE with additional options (e.g., compaction settings)
		// TODO: Unsupported compaction settings
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT) WITH compaction = {'class': 'LeveledCompactionStrategy'};",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
			}},
		},

		// CREATE MATERIALIZED VIEW
		// TODO: Unsupported materialized views
		// 		{
		// 			query:       "CREATE MATERIALIZED VIEW user_by_name AS SELECT id, name FROM users WHERE name IS NOT NULL PRIMARY KEY (name, id);",
		// 			expectedErr: false,
		// 		},

		// CREATE INDEX on a column
		// TODO: Unsupported indexes
		// 		{
		// 			query:       "CREATE INDEX ON users(name);",
		// 			expectedErr: false,
		// 		},

		// DROP TABLE
		// TODO: Unsupported drop tables
		// 		{
		// 			query:       "DROP TABLE users;",
		// 			expectedErr: false,
		// 		},

		// Invalid CQL - Missing Column Definition in CREATE TABLE
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, );",
			expectedErr: true,
		},
		// Invalid CQL - Missing Parenthesis
		{
			query:       "CREATE TABLE users id UUID PRIMARY KEY;",
			expectedErr: true,
		},
		// CREATE TABLE with custom data type
		{
			query:       `CREATE TABLE users (id UUID PRIMARY KEY, name CUSTOMTYPE);`,
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeCustom("customtype")},
					}},
				}},
			}},
		},
		// DROP INDEX
		// TODO: Unsupported drop index
		// 		{
		// 			query:       "DROP INDEX idx_name;",
		// 			expectedErr: false,
		// 		},
		// Invalid DROP TABLE (non-existent table)
		// TODO: Unsupported drop tables
		// 		{
		// 			query:       "DROP TABLE nonexistent_table;",
		// 			expectedErr: false,
		// 		},
		// CREATE TABLE with a primary key that uses multiple columns (composite key)
		{
			query:       "CREATE TABLE customer_orders (customer_id UUID, order_id UUID, product TEXT, PRIMARY KEY (customer_id, order_id));",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "customer_orders", Columns: []*Column{
						{Name: "customer_id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "order_id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "product", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
			}},
		},

		// CREATE TABLE with static column
		// TODO: Unsupported static columns
		//{
		//	query: "CREATE TABLE blog (blog_id UUID PRIMARY KEY, title TEXT, description TEXT STATIC, num_views COUNTER);",
		//},

		// CREATE TABLE with a complex column type (frozen set of tuples)
		// TODO: Unsupported tuples
		// {
		// 	query:       "CREATE TABLE orders (id UUID PRIMARY KEY, items FROZEN<SET<TUPLE<UUID, TEXT>>>);",
		// 	expectedErr: false,
		// 	expectedSchema: &Schema{Keyspaces: []*Keyspace{
		// 		{Name: defaultKeyspaceName, Tables: []*Table{
		// 			{Name: "orders", Columns: []*Column{
		// 				{Name: "id", DataType: NativeTypeUUID.IntoDataType()},
		// 				{Name: "items", DataType: FrozenType{DataType: CollectionTypeSet{}}},
		// 				{Name: "product", DataType: NativeTypeText.IntoDataType()},
		// 			}},
		// 		}},
		// 	}},
		// },

		// CREATE TABLE with frozen set of tuples and TTL
		// TODO: Unsupported tuples
		// 		{
		// 			query:       "CREATE TABLE orders (id UUID PRIMARY KEY, items FROZEN<SET<TUPLE<UUID, TEXT>>>) WITH TTL 3600;",
		// 			expectedErr: false,
		// 		},

		// CREATE TABLE with static columns
		// TODO: Unsupported reserved keyword "timestamp"
		// {
		// 	query:       "CREATE TABLE sensor_data (sensor_id UUID, timestamp TIMESTAMP, temperature DECIMAL, PRIMARY KEY (sensor_id, timestamp)) WITH STATIC temperature;",
		// 	expectedErr: false,
		// 	expectedSchema: &Schema{Keyspaces: []*Keyspace{
		// 		{Name: defaultKeyspaceName, Tables: []*Table{
		// 			{Name: "sensor_data", Columns: []*Column{
		// 				{Name: "sensor_id", DataType: NativeTypeUUID.IntoDataType()},
		// 				{Name: "product", DataType: NativeTypeText.IntoDataType()},
		// 			}},
		// 		}},
		// 	}},
		// },

		// CREATE TABLE with a compound clustering key and order
		{
			query:       "CREATE TABLE articles (author TEXT, category TEXT, published_at TIMESTAMP, title TEXT, PRIMARY KEY (author, category, published_at)) WITH CLUSTERING ORDER BY (published_at DESC);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "articles", Columns: []*Column{
						{Name: "author", DataType: gocqlhelpers.NewTypeText()},
						{Name: "category", DataType: gocqlhelpers.NewTypeText()},
						{Name: "published_at", DataType: gocqlhelpers.NewTypeTimestamp()},
						{Name: "title", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
			}},
		},

		// Invalid CREATE MATERIALIZED VIEW due to missing WHERE clause
		// TODO: Unsupported materialized views
		// 		{
		// 			query:       "CREATE MATERIALIZED VIEW mv_user_by_name AS SELECT id, name FROM users PRIMARY KEY (name, id);",
		// 			expectedErr: false,
		// 		},

		// CREATE TABLE with custom validation
		{
			query:       "CREATE TABLE users (id UUID PRIMARY KEY, name TEXT) WITH validation = {'validator': 'my_custom_validator'};",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
			}},
		},

		// CREATE TABLE with keyspace specifier
		// TODO: Unsupported keyspace specifier
		{
			query:       "CREATE TABLE auth.users (id UUID PRIMARY KEY, name TEXT);",
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: "auth", Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeText()},
					}},
				}},
				{Name: defaultKeyspaceName, Tables: []*Table{}},
			}},
		},

		{
			query: `
		CREATE TABLE users (id UUID PRIMARY KEY, name TEXT, age INT);
		CREATE TABLE orders (order_id UUID, customer_id UUID, total DECIMAL, PRIMARY KEY (order_id, customer_id));
		CREATE TABLE logins (id UUID PRIMARY KEY, last_seen TIMESTAMP);
		      `,
			expectedErr: false,
			expectedSchema: &Schema{Keyspaces: []*Keyspace{
				{Name: defaultKeyspaceName, Tables: []*Table{
					{Name: "users", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "name", DataType: gocqlhelpers.NewTypeText()},
						{Name: "age", DataType: gocqlhelpers.NewTypeInt()},
					}},
					{Name: "orders", Columns: []*Column{
						{Name: "order_id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "customer_id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "total", DataType: gocqlhelpers.NewTypeDecimal()},
					}},
					{Name: "logins", Columns: []*Column{
						{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "last_seen", DataType: gocqlhelpers.NewTypeTimestamp()},
					}},
				}},
			}},
		},
	}
	parser := NewSchemaParser()
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			schema, err := parser.Parse(tt.query)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, but got: %v for query: %s", tt.expectedErr, err, tt.query)
				fmt.Println(schema.String())
			}
			if err != nil || tt.expectedErr {
				return
			}
			require.NotNil(t, schema)
			require.Equal(t, tt.expectedSchema.String(), schema.String())
		})
	}
}
