package cqlc

type Query struct {
	FuncName    string
	Annotations []string
	Stmt        string
	Params      []string
	Selects     []string
}
