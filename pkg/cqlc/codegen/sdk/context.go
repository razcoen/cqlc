package sdk

type Context struct {
	Provider    *SchemaQueriesProvider
	SchemaPath  string
	QueriesPath string
	ConfigPath  string
	Version     string
}
