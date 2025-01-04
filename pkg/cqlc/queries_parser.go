package cqlc

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/internal/antlrcql"
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
			return nil, fmt.Errorf("invalid query: %s", query)
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
		antlr.ParseTreeWalkerDefault.Walk(l, p.Cql())
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
			Params:      l.params,
			Selects:     l.selects,
		})
	}
	return queries, nil
}

type queriesParserListener struct {
	baseCQLParserListener
	selects []string
	params  []string
	err     error
}

func newQueriesParserListener() *queriesParserListener {
	return &queriesParserListener{
		baseCQLParserListener: baseCQLParserListener{debug: true},
	}
}

func (l *queriesParserListener) EnterSelect_(c *antlrcql.Select_Context) {
	var selectElements *antlrcql.SelectElementsContext
	var fromSpec *antlrcql.FromSpecContext
	for _, child := range c.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.SelectElementsContext:
			selectElements = child
		case *antlrcql.FromSpecContext:
			fromSpec = child
		}
	}
	if selectElements == nil {
		fmt.Println("WTF1")
		// TODO
		return
	}
	if fromSpec == nil {
		fmt.Println("WTF2")
		// TODO
		return
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
}

func (l *queriesParserListener) EnterWhereSpec(c *antlrcql.WhereSpecContext) {
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
