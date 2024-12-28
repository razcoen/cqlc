package cql

type Schema struct {
	Keyspaces []*Keyspace
}

type Keyspace struct {
	Name             string
	Tables           []*Table
	UserDefinedTypes []*UserDefinedType
}

type Table struct {
	Name    string
	Columns []*Column
}

type Column struct {
	Name     string
	DataType *DataType
}
