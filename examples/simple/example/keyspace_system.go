// Code generated by cqlc. DO NOT EDIT.

package example

import (
	"github.com/gocql/gocql"
	"gopkg.in/inf.v0"
	"time"
)

// Table: users
type User struct {
	UserID    gocql.UUID
	Username  string
	Email     string
	CreatedAt time.Time
}

// Table: posts
type Post struct {
	PostID    gocql.UUID
	UserID    gocql.UUID
	Content   string
	CreatedAt time.Time
}

// Table: user_profiles
type UserProfile struct {
	UserID    gocql.UUID
	FirstName string
	LastName  string
	Interest  []string
}

// Table: orders
type Order struct {
	OrderID   gocql.UUID
	UserID    gocql.UUID
	ProductID gocql.UUID
	Quantity  int
	OrderDate time.Time
}

// Table: customer_addresses
type CustomerAddress struct {
	CustomerID gocql.UUID
}

// Table: products
type Product struct {
	ProductID gocql.UUID
	Name      string
	Category  string
	Price     *inf.Dec
}

// Table: temporary_sessions
type TemporarySession struct {
	SessionID gocql.UUID
	UserID    gocql.UUID
	CreatedAt time.Time
}
