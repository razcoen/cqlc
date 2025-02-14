package golang

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
	"github.com/razcoen/cqlc/pkg/log"
)

type Generator struct {
	queriesGoTemplate *template.Template
	execQueryTemplate *template.Template
	oneQueryTemplate  *template.Template
	clientTemplate    *template.Template
	logger            log.Logger
}

func NewGenerator(logger log.Logger) *Generator {
	return &Generator{
		queriesGoTemplate: template.Must(template.New("queries-template").Parse(queriesGoTemplate)),
		clientTemplate:    template.Must(template.New("client-template").Parse(clientTemplate)),
		execQueryTemplate: template.Must(template.New("exec-query-template").Parse(execQueryGoTemplate)),
		oneQueryTemplate:  template.Must(template.New("one-query-template").Parse(oneQueryGoTemplate)),
		logger:            logger,
	}
}

func (gen *Generator) Generate(ctx *sdk.Context, opts *Options) (err error) {
	dir, err := os.MkdirTemp("", "cqlc-gen-go-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	logger := gen.logger.With("workdir", dir)
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			logger.Warn("failed to remove temporary workdir")
		}
	}()
	var filenames []string
	filename := "client.go"
	path := filepath.Join(dir, filename)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("open client file: %w", err)
	}
	filenames = append(filenames, filename)
	defer func() {
		if err := f.Close(); err != nil {
			logger.Warn("error closing file", "filepath", path, "error", err)
		}
	}()
	if err := gen.generateClient(ctx, &generateClientRequest{options: opts, out: f}); err != nil {
		return fmt.Errorf("generate client: %w", err)
	}

	for _, k := range ctx.Provider.Schema().Keyspaces {
		imports := make(map[string]bool)
		structByTableName := make(map[string]*strct, len(k.Tables))
		for _, t := range k.Tables {
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
			structByTableName[t.Name] = &strct{
				name:              st.Name,
				fieldByColumnName: fieldByColumnName,
			}
		}
		for _, t := range k.Tables {
			err := func() error {
				queries := ctx.Provider.ListTableQueries(k.Name, t.Name)
				if len(queries) == 0 {
					return nil
				}
				path := sdk.ToSnakeCase(t.Name) + ".go"
				if k.Name != "" {
					path = sdk.ToSnakeCase(k.Name) + "_" + path
				}
				path = "query_" + path
				filename := path
				filenames = append(filenames, path)
				path = filepath.Join(dir, path)
				f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
				if err != nil {
					return fmt.Errorf("open queries file: %w", err)
				}
				defer func() {
					if err := f.Close(); err != nil {
						gen.logger.Error("error closing file", "filepath", path, "error", err)
					}
				}()
				if err := gen.generateQueries(ctx, &generateQueriesRequest{
					queries:           queries,
					structByTableName: structByTableName,
					options:           opts,
					out:               f,
					path:              filepath.Join(opts.Out, filename),
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

	// Given that the client is generated in a temporary file, we need to move it to the output directory.
	out := opts.Out
	if err := os.MkdirAll(out, 0777); err != nil && !os.IsExist(err) {
		return fmt.Errorf("create output directory: %w", err)
	}
	for _, f := range filenames {
		if err := os.Remove(filepath.Join(out, f)); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove file %s: %w", f, err)
		}
	}
	if err := os.CopyFS(out, os.DirFS(dir)); err != nil {
		return fmt.Errorf("copy client file: %w", err)
	}
	return nil
}

var (
	clientTemplate = `
package {{.PackageName}}

import (
	"fmt"
	"errors"
	"github.com/gocql/gocql"
  "github.com/razcoen/cqlc/pkg/log"
  "github.com/razcoen/cqlc/pkg/gocqlc"
)

type Client struct {
	gocqlc.Client
}

func NewClient(session *gocql.Session, opts ...gocqlc.ClientOption) (*Client, error) {
	client, err := gocqlc.NewClient(session, opts...)
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}
	return &Client{Client: *client}, nil
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

	queriesGoTemplate = `
package {{.PackageName}}

import (
{{- range .Imports}}
  "{{.}}"
{{- end}}
	"fmt"
	"context"
	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/log"
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
	logger log.Logger
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
	session := c.Session()
	q := session.Query("{{.Stmt}}"{{- range .Params -}}, params.{{.Name}}{{- end -}})
	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
	for _, opt := range opts {
		q = opt.Apply(q)
	}
	return &{{.FuncName}}Querier{query: q, logger: c.Logger()}
}

{{else}}
{{if eq "batch" .Annotation}}
func (c *Client) {{.FuncName}}(ctx context.Context{{if .ParamsType}}, params []*{{.ParamsType}}{{end}}, opts ...gocqlc.BatchOption) error {
	session := c.Session()
	b := session.NewBatch(gocql.{{$.BatchType}}Batch)
	for _, v := range params {
		b.Query("{{.Stmt}}"{{- range .Params -}}, v.{{.Name}}{{- end -}})
	}
	b = b.WithContext(ctx)
	for _, opt := range c.DefaultBatchOptions() {
		b = opt.Apply(b)
	}
	for _, opt := range opts {
		b = opt.Apply(b)
	}
	if err := session.ExecuteBatch(b); err != nil {
		return fmt.Errorf("exec batch: %w", err)
	}
	return nil
}
{{ else }}
func (c *Client) {{.FuncName}}(ctx context.Context{{if .ParamsType}}, params *{{.ParamsType}}{{end}}, opts ...gocqlc.QueryOption) {{- if .ResultType -}}(*{{.ResultType}}, error){{- else -}}error{{- end -}} {
	session := c.Session()
	q := session.Query("{{.Stmt}}"{{- range .Params -}}, params.{{.Name}}{{- end -}})
  	q = q.WithContext(ctx)
   	for _, opt := range c.DefaultQueryOptions() {
		q = opt.Apply(q)
	}
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

type queriesGoTemplateValue struct {
	PackageName string
	Imports     []string
	Queries     []queryGoTemplateValue
	BatchType   string
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
	options           *Options
	queries           []*sdk.Query
	structByTableName map[string]*strct
	out               io.Writer
	path              string
}

func (gg *Generator) generateQueries(ctx *sdk.Context, req *generateQueriesRequest) error {
	batchType := parseBatchType(req.options.Defaults.BatchType, gocql.UnloggedBatch, gg.logger)
	v := queriesGoTemplateValue{
		PackageName: req.options.Package,
		BatchType:   batchType,
	}
	imports := make(map[string]bool)
	for _, q := range req.queries {
		params := make([]fieldTemplateValue, 0, len(q.Params))
		selects := make([]fieldTemplateValue, 0, len(q.Selects))
		strct := req.structByTableName[q.Table]
		for _, p := range q.Params {
			field, ok := strct.fieldByColumnName[p]
			if !ok {
				return fmt.Errorf(`unfamiliar column "%s" found in query "%s"`, p, q.Stmt)
			}
			if field.goType.ImportPath != "" {
				imports[field.goType.ImportPath] = true
			}
			params = append(params, fieldTemplateValue{Name: field.name, GoType: field.goType.Name})
		}
		for _, s := range q.Selects {
			f := strct.fieldByColumnName[s]
			if f.goType.ImportPath != "" {
				imports[f.goType.ImportPath] = true
			}
			selects = append(selects, fieldTemplateValue{Name: f.name, GoType: f.goType.Name})
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
	_, _ = buf.WriteString(createHeader(ctx, req.options, gg.logger))
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
	options *Options
	out     io.Writer
}

type clientTemplateValue struct {
	PackageName string
}

func (gg *Generator) generateClient(ctx *sdk.Context, req *generateClientRequest) error {
	buf := &bytes.Buffer{}
	_, _ = buf.WriteString(createHeader(ctx, req.options, gg.logger))
	if err := gg.clientTemplate.Execute(buf, &clientTemplateValue{PackageName: req.options.Package}); err != nil {
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

func parseBatchType(input string, fallback gocql.BatchType, logger log.Logger) string {
	parsed := fallback
	mapping := map[gocql.BatchType]string{gocql.UnloggedBatch: "Unlogged", gocql.LoggedBatch: "Logged", gocql.CounterBatch: "Counter"}
	foundValidBatchType := false
	for k, str := range mapping {
		if strings.ToLower(str) == input {
			parsed = k
			foundValidBatchType = true
			break
		}
	}
	if len(input) > 0 && !foundValidBatchType && logger != nil {
		logger.Warn(fmt.Sprintf("using default batch type %q: invalid batch type %q was provided", strings.ToLower(mapping[parsed]), input))
	}
	return mapping[parsed]
}
