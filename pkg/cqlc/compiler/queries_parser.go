package compiler

import (
	"errors"
	"fmt"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler/internal/antlrcql"
)

type QueriesParser struct{}

func NewQueriesParser() *QueriesParser {
	return &QueriesParser{}
}

type queryStmt struct {
	cql      antlrcql.ICqlContext
	comments []string
	text     string
}

func (qp *QueriesParser) Parse(cql string) (sdk.Queries, error) {
	el := newErrorListener()
	lexer := antlrcql.NewCQLLexer(antlr.NewInputStream(cql))
	lexer.AddErrorListener(el)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := antlrcql.NewCQLParser(stream)
	p.AddErrorListener(el)
	cqls := p.Cqls().AllCql()
	if el.errors != nil {
		return nil, errors.Join(el.errors...)
	}

	lineNumber := 1
	var stmts []*queryStmt
	lines := strings.Split(cql, "\n")
	for _, cql := range cqls {
		stmt := &queryStmt{
			cql:  cql,
			text: strings.Join(lines[cql.GetStart().GetLine()-1:cql.GetStop().GetLine()], "\n"),
		}
		for i := lineNumber - 1; i < cql.GetStart().GetLine()-1; i++ {
			stmt.comments = append(stmt.comments, lines[i])
		}
		stmts = append(stmts, stmt)
		lineNumber = cql.GetStop().GetLine() + 1
	}

	queries := make(sdk.Queries, 0, len(stmts))
	for _, stmt := range stmts {
		l := newQueriesParserListener()
		var t antlr.Tree = stmt.cql
		t = t.GetChild(0)
		switch t := t.(type) {
		case *antlrcql.Select_Context:
			l.EnterSelect_(t)
		case *antlrcql.InsertContext:
			l.EnterInsert(t)
		case *antlrcql.Delete_Context:
			l.EnterDelete_(t)
		}
		if l.err != nil {
			return nil, fmt.Errorf("error during traversal: %w", l.err)
		}
		funcName, annotations, err := parseComment(stmt.comments[len(stmt.comments)-1])
		if err != nil {
			return nil, fmt.Errorf("error parsing comment: %w", err)
		}
		queries = append(queries, &sdk.Query{
			FuncName:    funcName,
			Annotations: annotations,
			Stmt:        stmt.text,
			Table:       l.table,
			Params:      l.params,
			Selects:     l.selects,
			Keyspace:    l.keyspace,
		})
	}
	return queries, nil
}

type queriesParserListener struct {
	antlrcql.BaseCQLParserListener
	selects  []string
	params   []string
	table    string
	keyspace string
	err      error
}

func newQueriesParserListener() *queriesParserListener {
	return &queriesParserListener{}
}

func (l *queriesParserListener) EnterInsert(ctx *antlrcql.InsertContext) {
	var insertColumnSpec *antlrcql.InsertColumnSpecContext
	var insertValuesSpec *antlrcql.InsertValuesSpecContext
	var table *antlrcql.TableContext
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.InsertColumnSpecContext:
			insertColumnSpec = child
		case *antlrcql.InsertValuesSpecContext:
			insertValuesSpec = child
		case *antlrcql.TableContext:
			table = child
		}
	}
	var columns []string
	var paramColumns []bool
	for _, child := range insertColumnSpec.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.ColumnListContext:
			for _, grandchild := range child.GetChildren() {
				switch grandchild := grandchild.(type) {
				case *antlrcql.ColumnContext:
					columns = append(columns, grandchild.GetText())
				}
			}
		}
	}
	for _, child := range insertValuesSpec.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.ExpressionListContext:
			for _, grandchild := range child.GetChildren() {
				switch grandchild := grandchild.(type) {
				case *antlrcql.ExpressionContext:
					for _, grandgrandchild := range grandchild.GetChildren() {
						switch grandgrandchild.(type) {
						case *antlrcql.ConstantContext:
							isParam := false
							if grandchild.GetText() == "?" {
								isParam = true
							}
							paramColumns = append(paramColumns, isParam)
						}
					}
				}
			}
		}
	}
	params := make([]string, 0, len(columns))
	for i, column := range columns {
		if paramColumns[i] {
			params = append(params, column)
		}
	}
	l.params = params
	l.table = table.GetText()
}

func (l *queriesParserListener) EnterSelect_(c *antlrcql.Select_Context) {
	var selectElements *antlrcql.SelectElementsContext
	var fromSpec *antlrcql.FromSpecContext
	var whereSpec *antlrcql.WhereSpecContext
	for _, child := range c.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.SelectElementsContext:
			selectElements = child
		case *antlrcql.FromSpecContext:
			fromSpec = child
		case *antlrcql.WhereSpecContext:
			whereSpec = child
		}
	}
	var selects []string
	for _, child := range selectElements.GetChildren() {
		switch child := child.(type) {
		case antlr.TerminalNode:
			selects = append(selects, child.GetText())
		case *antlrcql.SelectElementContext:
			for _, grandchild := range child.GetChildren() {
				switch grandchild := grandchild.(type) {
				case antlr.TerminalNode:
					selects = append(selects, grandchild.GetText())
				}
			}
		}
	}
	l.selects = selects
	l.parseWhereSpec(whereSpec)
	l.parseFromSpec(fromSpec)
}

func (l *queriesParserListener) parseFromSpec(c *antlrcql.FromSpecContext) {
	if c.GetChildCount() != 2 {
		return
	}
	if _, ok := c.GetChild(0).(*antlrcql.KwFromContext); !ok {
		return
	}
	fromSpecElement, ok := c.GetChild(1).(*antlrcql.FromSpecElementContext)
	if !ok {
		return
	}
	if fromSpecElement.GetChildCount() == 1 {
		tn, ok := fromSpecElement.GetChild(0).(antlr.TerminalNode)
		if !ok {
			return
		}
		l.table = tn.GetText()
	}
	if fromSpecElement.GetChildCount() == 3 {
		dot, ok := fromSpecElement.GetChild(1).(antlr.TerminalNode)
		if !ok || dot.GetText() != "." {
			return
		}
		kn, ok := fromSpecElement.GetChild(0).(antlr.TerminalNode)
		if !ok {
			return
		}
		tn, ok := fromSpecElement.GetChild(2).(antlr.TerminalNode)
		if !ok {
			return
		}
		l.table = tn.GetText()
		l.keyspace = kn.GetText()
	}
}

func (l *queriesParserListener) parseWhereSpec(c *antlrcql.WhereSpecContext) {
	if c == nil {
		return
	}
	var relationElements *antlrcql.RelationElementsContext
	for _, child := range c.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.RelationElementsContext:
			relationElements = child
		}
	}
	if relationElements == nil {
		fmt.Println("WTF3")
		// TODO
		return
	}
	var params []string
	for _, child := range relationElements.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.RelationElementContext:
			var isParam bool
			var param string
			for _, grandchild := range child.GetChildren() {
				switch grandchild := grandchild.(type) {
				case antlr.TerminalNode:
					validColumn := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
					if validColumn.MatchString(grandchild.GetText()) {
						param = grandchild.GetText()
					}
				case *antlrcql.ConstantContext:
					if grandchild.GetText() == "?" {
						isParam = true
					}
				}
			}
			if isParam && param != "" {
				params = append(params, param)
			}
		}
	}
	l.params = params
}

func (l *queriesParserListener) EnterDelete_(c *antlrcql.Delete_Context) {
	var fromSpec *antlrcql.FromSpecContext
	var whereSpec *antlrcql.WhereSpecContext
	for _, child := range c.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.FromSpecContext:
			fromSpec = child
		case *antlrcql.WhereSpecContext:
			whereSpec = child
		}
	}
	if fromSpec == nil || whereSpec == nil {
		return
	}
	l.parseFromSpec(fromSpec)
	l.parseWhereSpec(whereSpec)
}

func parseComment(comment string) (string, []string, error) {
	regex := regexp.MustCompile(`--\s*name:\s*(\w+)\s*(:\w+)*`)
	matches := regex.FindStringSubmatch(comment)
	if matches == nil || len(matches) < 2 {
		return "", nil, fmt.Errorf("invalid comment structure")
	}
	funcName := matches[1]
	annotations := strings.Fields(matches[2])
	for i, ann := range annotations {
		annotations[i] = strings.TrimPrefix(ann, ":")
	}
	return funcName, annotations, nil
}
