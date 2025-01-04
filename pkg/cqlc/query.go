package cqlc

type Query struct {
	Stmt    string
	Params  []string
	Selects []string
}
