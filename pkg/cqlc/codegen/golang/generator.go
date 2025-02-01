package golang

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"text/template"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
	"github.com/razcoen/cqlc/pkg/cqlc/log"
)

type Generator struct {
	keyspaceGoTemplate *template.Template
	queriesGoTemplate  *template.Template
	execQueryTemplate  *template.Template
	oneQueryTemplate   *template.Template
	clientTemplate     *template.Template
	logger             log.Logger
}

func NewGenerator(logger log.Logger) (*Generator, error) {
	tpl, err := template.New("keyspace-go-template").Parse(keyspaceGoTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse keyspace template: %w", err)
	}
	tpl2, err := template.New("queries-go-template").Parse(queriesGoTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse queries template: %w", err)
	}
	tpl3, err := template.New("client-template").Parse(clientTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse client template: %w", err)
	}
	tpl4, err := template.New("exec-query-template").Parse(execQueryGoTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse exec query template: %w", err)
	}
	tpl5, err := template.New("one-query-template").Parse(oneQueryGoTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse one query template: %w", err)
	}
	return &Generator{
		keyspaceGoTemplate: tpl,
		queriesGoTemplate:  tpl2,
		clientTemplate:     tpl3,
		execQueryTemplate:  tpl4,
		oneQueryTemplate:   tpl5,
		logger:             logger,
	}, nil
}

func (gg *Generator) Generate(req *sdk.GenerateRequest, opts *Options) error {
	schema := req.Schema
	queries := req.Queries
	out := opts.Out
	if err := os.MkdirAll(out, 0777); err != nil && !os.IsExist(err) {
		return fmt.Errorf("create output directory: %w", err)
	}
	fn := filepath.Join(out, "client.go")
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("open client file: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			gg.logger.Error("error closing file", "filepath", fn, "error", err)
		}
	}()
	if err := gg.generateClient(&generateClientRequest{
		packageName: opts.Package,
		out:         f,
	}); err != nil {
		return fmt.Errorf("generate client: %w", err)
	}

	keyspaceTableSet := make(map[string]map[string]bool)
	for _, k := range schema.Keyspaces {
		if _, ok := keyspaceTableSet[k.Name]; !ok {
			keyspaceTableSet[k.Name] = make(map[string]bool)
		}
		for _, t := range k.Tables {
			keyspaceTableSet[k.Name][t.Name] = true
		}
	}

	queriesByTableByKeyspace := make(map[string]map[string]sdk.Queries)
	var invalidQueriesErrs []error
	for _, q := range queries {
		// Validate that the query keyspace and table exist in the schema.
		if _, ok := keyspaceTableSet[q.Keyspace]; !ok {
			invalidQueriesErrs = append(invalidQueriesErrs, fmt.Errorf("keyspace %q does not exist in schema", q.Keyspace))
		}
		if _, ok := keyspaceTableSet[q.Keyspace][q.Table]; !ok {
			invalidQueriesErrs = append(invalidQueriesErrs, fmt.Errorf("table %q does not exist in keyspace %q", q.Table, q.Keyspace))
		}
		// Map queries by keyspace and table.
		if _, ok := queriesByTableByKeyspace[q.Keyspace]; !ok {
			queriesByTableByKeyspace[q.Keyspace] = make(map[string]sdk.Queries)
		}
		if _, ok := queriesByTableByKeyspace[q.Keyspace][q.Table]; !ok {
			queriesByTableByKeyspace[q.Keyspace][q.Table] = make(sdk.Queries, 0)
		}
		queriesByTableByKeyspace[q.Keyspace][q.Table] = append(queriesByTableByKeyspace[q.Keyspace][q.Table], q)
	}

	if len(invalidQueriesErrs) > 0 {
		return fmt.Errorf("invalid queries: %w", errors.Join(invalidQueriesErrs...))
	}

	for _, k := range schema.Keyspaces {
		resp, err := gg.generateKeyspaceStructs(&generateKeyspaceStructsRequest{
			keyspace:    k,
			packageName: opts.Package,
			out:         noopWriter{},
		})
		if err != nil {
			return fmt.Errorf("generate keyspace: %w", err)
		}
		for _, t := range k.Tables {
			err := func() error {
				queries := queriesByTableByKeyspace[k.Name][t.Name]
				if len(queries) == 0 {
					return nil
				}
				fn := sdk.ToSnakeCase(t.Name) + ".go"
				if k.Name != "" {
					fn = sdk.ToSnakeCase(k.Name) + "_" + fn
				}
				fn = "query_" + fn
				fn = filepath.Join(out, fn)
				f, err := os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
				if err != nil {
					return fmt.Errorf("open queries file: %w", err)
				}
				defer func() {
					if err := f.Close(); err != nil {
						gg.logger.Error("error closing file", "filepath", fn, "error", err)
					}
				}()
				if err := gg.generateQueries(&generateQueriesRequest{
					queries:           queries,
					structByTableName: resp.structByTableName,
					packageName:       opts.Package,
					out:               f,
				}); err != nil {
					return fmt.Errorf("generate queries: %w", err)
				}
				return nil
			}()
			if err != nil {
				return fmt.Errorf("generate queries: %w", err)
			}
		}
	}
	return nil
}

type generateKeyspaceStructsRequest struct {
	keyspace    *sdk.Keyspace
	packageName string
	out         io.Writer
}

var (
	// TODO: Create a header that is always appended
	keyspaceGoTemplate = `// Code generated by cqlc. DO NOT EDIT.

package {{.PackageName}}

{{- if gt (len .Imports) 0}}
import (
{{- end}}
{{- range .Imports}}
  "{{.}}"
{{- end}}
{{- if gt (len .Imports) 0}}
)
{{- end}}
{{range .Structs}}
// Table: {{.TableName}}
type {{.Name}} struct {
  {{- range .Fields }}
  {{ .Name }} {{ .GoType }}
  {{- end }}
}
{{end -}}
`

	clientTemplate = `// Code generated by cqlc. DO NOT EDIT.

package {{.PackageName}}

import (
	"fmt"
	"github.com/gocql/gocql"
  "github.com/razcoen/cqlc/pkg/gocqlc"
)

type Client struct {
	session *gocql.Session
  logger gocqlc.Logger
}

func NewClient(session *gocql.Session, logger gocqlc.Logger) (*Client, error) {
	if session == nil {
		return nil, fmt.Errorf("session cannot be nil")
	}
	if session.Closed() {
		return nil, fmt.Errorf("session already closed")
	}
	if logger == nil {
		logger = &gocqlc.NoopLogger{}
	}
	return &Client{session: session, logger: logger}, nil
}

func (c *Client) Close() error {
	c.session.Close()
	return nil
}
`
	oneQueryGoTemplate = `
	var result {{.ResultType}}
	if err := q.Scan({{- range .Selects -}}&result.{{.Name}},{{- end -}}); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &result, nil`
	execQueryGoTemplate = `
	if err := q.Exec(); err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	return nil`

	queriesGoTemplate = `// Code generated by cqlc. DO NOT EDIT.

package {{.PackageName}}

import (
{{- range .Imports}}
  "{{.}}"
{{- end}}
	"fmt"
	"context"
	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/gocqlc"
)

{{range .Queries}}
{{if .ParamsType -}}
type {{.ParamsType}} struct {
{{- range .Params}}
{{.Name}} {{.GoType}}
{{- end}}
}
{{- end}}

{{if .ResultType -}}
type {{.ResultType}} struct {
{{- range .Selects}}
{{.Name}} {{.GoType}}
{{- end}}
}
{{- end}}

{{if eq "many" .Annotation}}
type {{.FuncName}}Querier struct {
	query *gocql.Query
	logger gocqlc.Logger
}

func (q *{{.FuncName}}Querier) All(ctx context.Context) ([]*{{.ResultType}}, error) {
	var results []*{{.ResultType}}
	var pageState []byte
	for {
		page, err := q.Page(ctx, pageState)
		if err != nil {
			return nil, fmt.Errorf("page: %w", err)
		}
		results = append(results, page.Results()...)
		if len(page.PageState()) == 0 {
			break
		}
		pageState = page.PageState()
	}
	return results, nil
}

type {{.ResultType}}sPage struct {
	results []*{{.ResultType}}
	pageState []byte
	numRows int
}

func (page *{{.ResultType}}sPage) Results() []*{{.ResultType}} { return page.results }
func (page *{{.ResultType}}sPage) NumRows() int { return page.numRows }
func (page *{{.ResultType}}sPage) PageState() []byte { return page.pageState }

func (q *{{.FuncName}}Querier) Page(ctx context.Context, pageState []byte) (*{{.ResultType}}sPage, error) {
	var results []*{{.ResultType}}
	iter := q.query.WithContext(ctx).PageState(pageState).Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			q.logger.Error("iter.Close() returned with error", "error", err)
		}
	} ()
	nextPageState := iter.PageState()
	scanner := iter.Scanner()
	for scanner.Next() {
		var result {{.ResultType}}
		if err := scanner.Scan({{- range .Selects -}}&result.{{.Name}},{{- end -}}); err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		results = append(results, &result)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}
	return &{{.ResultType}}sPage{results: results, pageState: nextPageState, numRows: iter.NumRows()}, nil
}

func (c *Client) {{.FuncName}}({{if .ParamsType}}params *{{.ParamsType}}, {{end}}opts ...gocqlc.QueryOption) *{{.FuncName}}Querier {
	q := c.session.Query("{{.Stmt}}"{{- range .Params -}}, params.{{.Name}}{{- end -}})
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	return &{{.FuncName}}Querier{query: q, logger: c.logger}
}

{{else}}
{{if eq "batch" .Annotation}}
func (c *Client) {{.FuncName}}(ctx context.Context{{if .ParamsType}}, params []*{{.ParamsType}}{{end}}, opts ...gocqlc.BatchOption) error {
	b := c.session.NewBatch(gocql.UnloggedBatch)
	for _, v := range params {
		b.Query("{{.Stmt}}"{{- range .Params -}}, v.{{.Name}}{{- end -}})
	}
  b = b.WithContext(ctx)
	for _, opt := range opts {
		b = opt.Apply(b)
	}
	if err := c.session.ExecuteBatch(b); err != nil {
		return fmt.Errorf("exec batch: %w", err)
	}
	return nil
}
{{ else }}
func (c *Client) {{.FuncName}}(ctx context.Context{{if .ParamsType}}, params *{{.ParamsType}}{{end}}, opts ...gocqlc.QueryOption) {{- if .ResultType -}}(*{{.ResultType}}, error){{- else -}}error{{- end -}} {
	q := c.session.Query("{{.Stmt}}"{{- range .Params -}}, params.{{.Name}}{{- end -}})
  q = q.WithContext(ctx)
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	{{- .ExecString}}
}
{{ end -}}
{{ end -}}

{{end -}}
`
)

// TODO: When returning "*" it might be needed to iterate through the column informations of the iterator just to be sure on the ordering.

type queriesGoTemplateValue struct {
	PackageName string
	Imports     []string
	Queries     []queryGoTemplateValue
}

type queryGoTemplateValue struct {
	ExecString string
	Annotation sdk.Annotation
	ParamsType string
	ResultType string
	FuncName   string
	Stmt       string
	Params     []fieldTemplateValue
	Selects    []fieldTemplateValue
}

type fieldTemplateValue struct {
	Name   string
	GoType string
}

type keyspaceGoTemplateValue struct {
	PackageName string
	Imports     []string
	Structs     []struct {
		TableName string
		Name      string
		Fields    []fieldTemplateValue
	}
}

type generateKeyspaceStructsResponse struct {
	structByTableName map[string]*strct
}

func (gg *Generator) generateKeyspaceStructs(req *generateKeyspaceStructsRequest) (*generateKeyspaceStructsResponse, error) {
	v := keyspaceGoTemplateValue{
		PackageName: req.packageName,
		Structs: []struct {
			TableName string
			Name      string
			Fields    []fieldTemplateValue
		}{},
	}
	imports := make(map[string]bool)
	structByTableName := make(map[string]*strct, len(req.keyspace.Tables))
	for _, t := range req.keyspace.Tables {
		structName := sdk.ToSingularPascalCase(t.Name)
		st := struct {
			TableName string
			Name      string
			Fields    []fieldTemplateValue
		}{
			TableName: t.Name,
			Name:      structName,
		}
		fieldByColumnName := make(map[string]*field)
		for i, c := range t.Columns {
			name := sdk.ToSingularPascalCase(c.Name)
			goType, err := gocqlhelpers.ParseGoType(c.DataType)
			if err != nil {
				// TODO
				continue
			}
			if goType.ImportPath != "" {
				imports[goType.ImportPath] = true
			}
			st.Fields = append(st.Fields, fieldTemplateValue{Name: name, GoType: goType.Name})
			fieldByColumnName[c.Name] = &field{name: name, goType: goType, ordering: i + 1}
		}
		v.Structs = append(v.Structs, st)
		structByTableName[t.Name] = &strct{
			name:              st.Name,
			fieldByColumnName: fieldByColumnName,
		}
	}
	v.Imports = slices.Collect(maps.Keys(imports))
	buf := &bytes.Buffer{}
	if err := gg.keyspaceGoTemplate.Execute(buf, v); err != nil {
		return nil, fmt.Errorf("execute keyspace template: %w", err)
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format out: %w", err)
	}
	if _, err := req.out.Write(out); err != nil {
		return nil, fmt.Errorf("write out: %w", err)
	}
	return &generateKeyspaceStructsResponse{structByTableName: structByTableName}, nil
}

type field struct {
	name     string
	goType   *gocqlhelpers.GoType
	ordering int
}

type strct struct {
	name              string
	fieldByColumnName map[string]*field
}

type generateQueriesRequest struct {
	queries           []*sdk.Query
	structByTableName map[string]*strct
	packageName       string
	out               io.Writer
}

// TODO: Support keyspaces
func (gg *Generator) generateQueries(req *generateQueriesRequest) error {
	v := queriesGoTemplateValue{
		PackageName: req.packageName,
	}
	imports := make(map[string]bool)
	for _, q := range req.queries {
		params := make([]fieldTemplateValue, 0, len(q.Params))
		selects := make([]fieldTemplateValue, 0, len(q.Selects))
		strct := req.structByTableName[q.Table]
		for _, p := range q.Params {
			field := strct.fieldByColumnName[p]
			if field.goType.ImportPath != "" {
				imports[field.goType.ImportPath] = true
			}
			params = append(params, fieldTemplateValue{Name: field.name, GoType: field.goType.Name})
		}
		// TODO: Return the struct instead of copying the fields
		if len(q.Selects) == 1 && q.Selects[0] == "*" {
			fields := slices.Collect(maps.Values(strct.fieldByColumnName))
			slices.SortFunc(fields, func(a, b *field) int { return a.ordering - b.ordering })
			for _, f := range fields {
				if f.goType.ImportPath != "" {
					imports[f.goType.ImportPath] = true
				}
				selects = append(selects, fieldTemplateValue{Name: f.name, GoType: f.goType.Name})
			}
		} else {
			for _, s := range q.Selects {
				f := strct.fieldByColumnName[s]
				if f.goType.ImportPath != "" {
					imports[f.goType.ImportPath] = true
				}
				selects = append(selects, fieldTemplateValue{Name: f.name, GoType: f.goType.Name})
			}
		}
		query := queryGoTemplateValue{FuncName: q.FuncName, Params: params, Selects: selects, Stmt: q.Stmt}
		var annotation sdk.Annotation
		for _, a := range q.Annotations {
			if qt, ok := sdk.ParseAnnotation(a); ok {
				annotation = qt
			}
		}
		query.Annotation = annotation
		// Set result type
		if (annotation == sdk.AnnotationOne || annotation == sdk.AnnotationMany) && len(query.Selects) > 0 {
			query.ResultType = fmt.Sprintf("%sResult", query.FuncName)
		}
		// Set params type
		if len(query.Params) > 0 {
			query.ParamsType = fmt.Sprintf("%sParams", query.FuncName)
		}
		switch annotation {
		case sdk.AnnotationExec:
			buf := &bytes.Buffer{}
			if err := gg.execQueryTemplate.Execute(buf, query); err != nil {
				return fmt.Errorf("execute exec query template: %w", err)
			}
			query.ExecString = buf.String()
		case sdk.AnnotationOne:
			buf := &bytes.Buffer{}
			if err := gg.oneQueryTemplate.Execute(buf, query); err != nil {
				return fmt.Errorf("execute one query template: %w", err)
			}
			query.ExecString = buf.String()
		case sdk.AnnotationBatch:

		}
		v.Queries = append(v.Queries, query)
	}
	v.Imports = slices.Collect(maps.Keys(imports))
	buf := &bytes.Buffer{}
	if err := gg.queriesGoTemplate.Execute(buf, v); err != nil {
		return fmt.Errorf("execute queries template: %w", err)
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return fmt.Errorf("format out: %w", err)
	}
	if _, err := req.out.Write(out); err != nil {
		return fmt.Errorf("write out: %w", err)
	}
	return nil
}

type generateClientRequest struct {
	packageName string
	out         io.Writer
}

type clientTemplateValue struct {
	PackageName string
}

func (gg *Generator) generateClient(req *generateClientRequest) error {
	buf := &bytes.Buffer{}
	if err := gg.clientTemplate.Execute(buf, &clientTemplateValue{PackageName: req.packageName}); err != nil {
		return fmt.Errorf("execute queries template: %w", err)
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format out: %w", err)
	}
	if _, err := req.out.Write(out); err != nil {
		return fmt.Errorf("write out: %w", err)
	}
	return nil
}

type noopWriter struct{}

func (w noopWriter) Write(b []byte) (n int, err error) { return 0, nil }
