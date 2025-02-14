package sdk

type Context struct {
	Provider *SchemaQueriesProvider
	Metadata *Metadata
}

type Metadata struct {
	SchemaPath  string
	QueriesPath string
	ConfigPath  string
	Version     string
}
