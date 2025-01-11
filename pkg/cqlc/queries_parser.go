package cqlc

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/internal/antlrcql"
	"github.com/xlab/treeprint"
	"reflect"
	"regexp"
	"strings"
)

type QueriesParser struct{}

func NewQueriesParser() *QueriesParser {
	return &QueriesParser{}
}

func (qp *QueriesParser) Parse(cql string) ([]*Query, error) {
	queryStmts := strings.Split(cql, ";")
	queries := make([]*Query, 0, len(queryStmts))
	for _, query := range queryStmts {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		items := strings.Split(query, "\n")
		if len(items) != 2 {
			return nil, fmt.Errorf("invalid query missing comment: %s", query)
		}
		comment := strings.TrimSpace(items[0])
		stmt := strings.TrimSpace(items[1])
		if !strings.HasPrefix(comment, "--") {
			return nil, fmt.Errorf("invalid query expected a comment: %s", query)
		}
		el := newErrorListener()
		l := newQueriesParserListener()
		lexer := antlrcql.NewCQLLexer(antlr.NewInputStream(stmt))
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(el)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := antlrcql.NewCQLParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(el)
		var t antlr.Tree = p.Cql()
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
		if el.errors != nil {
			return nil, errors.Join(el.errors...)
		}
		funcName, annotations, err := parseComment(comment)
		if err != nil {
			return nil, fmt.Errorf("error parsing comment: %w", err)
		}
		queries = append(queries, &Query{
			FuncName:    funcName,
			Annotations: annotations,
			Stmt:        stmt,
			Table:       l.table,
			Params:      l.params,
			Selects:     l.selects,
		})
	}
	return queries, nil
}

func printTree(t antlr.Tree) {
	p := treeprint.New()
	visit(p, t)
	fmt.Println(p.String())
}

func visit(p treeprint.Tree, t antlr.Tree) {
	type GetTexter interface {
		GetText() string
	}
	getTexter, ok := t.(GetTexter)
	if !ok {
		return
	}
	p1 := p.AddBranch(strings.Join([]string{reflect.TypeOf(t).String(), " :: ", getTexter.GetText()}, ""))
	for _, c := range t.GetChildren() {
		visit(p1, c)
	}
}

type queriesParserListener struct {
	antlrcql.BaseCQLParserListener
	selects []string
	params  []string
	table   string
	err     error
}

func newQueriesParserListener() *queriesParserListener {
	return &queriesParserListener{
		//baseCQLParserListener: baseCQLParserListener{},
	}
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
	if fromSpecElement.GetChildCount() != 1 {
		return
	}
	tn, ok := fromSpecElement.GetChild(0).(antlr.TerminalNode)
	if !ok {
		return
	}
	l.table = tn.GetText()
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
