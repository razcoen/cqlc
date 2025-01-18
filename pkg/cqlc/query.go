package cqlc

import "errors"

type Queries []*Query

type Query struct {
	FuncName    string
	Annotations []string
	Stmt        string
	Params      []string
	Selects     []string
	Table       string
	Keyspace    string
}

type QueryType string

const (
	QueryTypeExec QueryType = "exec"
	QueryTypeOne  QueryType = "one"
	QueryTypeMany QueryType = "many"
)

var errInvalidQueryType = errors.New("invalid query type")

func parseQueryType(s string) (QueryType, bool) {
	m := map[QueryType]bool{QueryTypeExec: true, QueryTypeOne: true, QueryTypeMany: true}
	qt := QueryType(s)
	_, ok := m[qt]
	return qt, ok
}
