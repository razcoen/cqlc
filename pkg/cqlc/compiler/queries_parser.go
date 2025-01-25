package compiler

import (
	"fmt"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler/internal/antlrcql"
	"regexp"
	"strings"
)

type QueriesParser struct{}

func NewQueriesParser() *QueriesParser {
	return &QueriesParser{}
}

func (qp *QueriesParser) Parse(cql string) (sdk.Queries, error) {
	cqls, err := antlrParse(cql)
	if err != nil {
		return nil, fmt.Errorf("antlr parse cql: %w", err)
	}

	type parsedQueryStmt struct {
		cql      antlrcql.ICqlContext
		comments []string
		text     string
	}
	var stmts []*parsedQueryStmt
	lineNumber := 1
	lines := strings.Split(cql, "\n")
	for _, cql := range cqls {
		// Join the astQueryAnalysis statement text into one line and remove all extra spaces.
		text := strings.Join(lines[cql.GetStart().GetLine()-1:cql.GetStop().GetLine()], " ")
		for strings.Contains(text, "  ") {
			text = strings.Replace(text, "  ", " ", -1)
		}
		stmt := &parsedQueryStmt{cql: cql, text: text}
		for i := lineNumber - 1; i < cql.GetStart().GetLine()-1; i++ {
			line := lines[i]
			if len(strings.TrimSpace(line)) == 0 {
				continue
			}
			stmt.comments = append(stmt.comments, line)
		}
		stmts = append(stmts, stmt)
		lineNumber = cql.GetStop().GetLine() + 1
	}

	queries := make(sdk.Queries, 0, len(stmts))
	for _, stmt := range stmts {
		visitor := newQueryAnalyzer()
		q := visitor.analyze(stmt.cql)
		funcName, annotations, err := parseComment(stmt.comments[len(stmt.comments)-1])
		if err != nil {
			return nil, fmt.Errorf("error parsing comment: %w", err)
		}
		queries = append(queries, &sdk.Query{
			FuncName:    funcName,
			Annotations: annotations,
			Stmt:        stmt.text,
			Table:       q.table,
			Params:      q.params,
			Selects:     q.selects,
			Keyspace:    q.keyspace,
		})
	}
	return queries, nil
}

func parseComment(comment string) (funcName string, annotations []string, err error) {
	regex := regexp.MustCompile(`--\s*name:\s*(\w+)\s*(:\w+)*`)
	matches := regex.FindStringSubmatch(comment)
	if matches == nil || len(matches) < 2 {
		return "", nil, fmt.Errorf("invalid comment structure")
	}
	funcName = matches[1]
	annotations = strings.Fields(matches[2])
	for i, ann := range annotations {
		annotations[i] = strings.TrimPrefix(ann, ":")
	}
	return funcName, annotations, nil
}
