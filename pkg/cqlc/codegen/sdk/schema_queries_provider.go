package sdk

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
)

type SchemaQueriesProvider struct {
	schema  *Schema
	queries Queries

	// keyspace --> table --> column
	schemaMap                map[string]map[string]*parsedTable
	queriesByTableByKeyspace map[string]map[string]Queries
}

// CompileSchemaWithQueries validates the given queries against the schema, and creates a provider.
// A provider exposes functions to ask questions about the queries that are being generated on the schema.
func CompileSchemaWithQueries(schema *Schema, queries Queries) (*SchemaQueriesProvider, error) {
	// TODO: Add keyspace inference option, meaning that the keyspace is not validated and tables are matched in a flat manner.
	functionNameValidator, err := regexp.Compile("^[A-Za-z_][A-Za-z0-9_:<>~]*$")
	if err != nil {
		return nil, fmt.Errorf("compile function name validation regexp: %w", err)
	}
	schemaMap := make(map[string]map[string]*parsedTable)
	for _, k := range schema.Keyspaces {
		if _, ok := schemaMap[k.Name]; !ok {
			schemaMap[k.Name] = make(map[string]*parsedTable)
		}
		keyspaceMap := schemaMap[k.Name]
		for _, t := range k.Tables {
			keyspaceMap[t.Name] = parseTable(t)
		}
	}
	queriesByTableByKeyspace := make(map[string]map[string]Queries)
	var invalidQueriesErrs []error
	for _, q := range queries {
		recordError := func(err error) {
			invalidQueriesErrs = append(invalidQueriesErrs, fmt.Errorf("query %q: %w", q.FuncName, err))
		}
		if !functionNameValidator.MatchString(q.FuncName) {
			recordError(fmt.Errorf("invalid query function name selected %q: string must follow the regexp %q", q.FuncName, functionNameValidator.String()))
			continue
		}
		if len(q.Annotations) == 0 {
			recordError(fmt.Errorf("missing annotations: use of of the following: %v", Annotations()))
			continue
		}
		// Validate that the query keyspace and table exist in the schema.
		if _, ok := schemaMap[q.Keyspace]; !ok {
			recordError(fmt.Errorf("keyspace %q does not exist in schema", q.Keyspace))
			continue
		}
		if _, ok := schemaMap[q.Keyspace][q.Table]; !ok {
			recordError(fmt.Errorf("table %q does not exist in keyspace %q", q.Table, q.Keyspace))
			continue
		}
		// Validate that all the query params are valid table columns.
		table := schemaMap[q.Keyspace][q.Table]
		formattedTableName := q.Table
		if q.Keyspace != "" {
			formattedTableName = fmt.Sprintf("%s.%s", q.Keyspace, q.Table)
		}
		validParams := true
		for _, c := range q.Params {
			if _, ok := table.columnByName[c]; !ok {
				recordError(fmt.Errorf("parametrized column %q does not exist in table %q", c, formattedTableName))
				validParams = false
			}
		}
		if !validParams {
			continue
		}
		selects := q.Selects
		if slices.Contains(q.Selects, "*") && len(q.Selects) > 1 {
			recordError(fmt.Errorf("cannot select both * and other columns: choose either * or specific columns"))
			continue
		}
		if slices.Contains(q.Selects, "*") && len(q.Selects) == 1 {
			selects = table.orderedColumns
		}
		validSelects := true
		for _, c := range selects {
			if _, ok := table.columnByName[c]; !ok {
				recordError(fmt.Errorf("selected column %q does not exist in table %q", c, formattedTableName))
				validSelects = false
			}
		}
		if !validSelects {
			continue
		}
		q := &Query{
			FuncName:    q.FuncName,
			Annotations: q.Annotations,
			Stmt:        q.Stmt,
			Params:      q.Params,
			// Overriding the selections to support "*" selection.
			Selects:  selects,
			Table:    q.Table,
			Keyspace: q.Keyspace,
		}
		// Map queries by keyspace and table.
		if _, ok := queriesByTableByKeyspace[q.Keyspace]; !ok {
			queriesByTableByKeyspace[q.Keyspace] = make(map[string]Queries)
		}
		if _, ok := queriesByTableByKeyspace[q.Keyspace][q.Table]; !ok {
			queriesByTableByKeyspace[q.Keyspace][q.Table] = make(Queries, 0)
		}
		queriesByTableByKeyspace[q.Keyspace][q.Table] = append(queriesByTableByKeyspace[q.Keyspace][q.Table], q)
	}
	err = nil
	if len(invalidQueriesErrs) > 0 {
		err = errors.Join(invalidQueriesErrs...)
	}
	return &SchemaQueriesProvider{
		schema:                   schema,
		queries:                  queries,
		schemaMap:                schemaMap,
		queriesByTableByKeyspace: queriesByTableByKeyspace,
	}, err
}

func (p *SchemaQueriesProvider) Schema() *Schema { return p.schema }

func (p *SchemaQueriesProvider) HasTable(keyspace, table string) bool {
	tables, keyspaceExists := p.schemaMap[keyspace]
	if !keyspaceExists {
		return false
	}
	_, tableExists := tables[table]
	return tableExists
}

func (p *SchemaQueriesProvider) ListTableQueries(keyspace, table string) Queries {
	tables, ok := p.queriesByTableByKeyspace[keyspace]
	if !ok {
		return nil
	}
	return tables[table]
}

type parsedTable struct {
	raw            *Table
	columnByName   map[string]*parsedColumn
	orderedColumns []string
}

type parsedColumn struct {
	*Column
	isPartOfPrimaryKey bool
}

func parseTable(t *Table) *parsedTable {
	ordering := make([]string, 0, len(t.Columns))
	primaryKeyColumns := make(map[string]bool)
	for _, c := range t.PrimaryKey.PartitionKey {
		primaryKeyColumns[c] = true
		ordering = append(ordering, c)
	}
	for _, c := range t.PrimaryKey.ClusteringKey {
		primaryKeyColumns[c] = true
		ordering = append(ordering, c)
	}
	var nonKeyColumns []string
	columnByName := make(map[string]*parsedColumn, len(t.Columns))
	for _, c := range t.Columns {
		isPartOfPrimaryKey := primaryKeyColumns[c.Name]
		columnByName[c.Name] = &parsedColumn{
			Column:             c,
			isPartOfPrimaryKey: isPartOfPrimaryKey,
		}
		if !isPartOfPrimaryKey {
			nonKeyColumns = append(nonKeyColumns, c.Name)
		}
	}
	slices.Sort(nonKeyColumns)
	ordering = append(ordering, nonKeyColumns...)
	return &parsedTable{raw: t, columnByName: columnByName, orderedColumns: ordering}
}
