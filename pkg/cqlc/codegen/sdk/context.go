package sdk

type Context struct {
	Schema      *Schema
	Queries     Queries
	SchemaPath  string
	QueriesPath string
	ConfigPath  string
	Version     string
}
