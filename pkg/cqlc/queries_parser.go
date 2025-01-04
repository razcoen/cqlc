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
	stmts := strings.Split(cql, ";")
	queries := make([]*Query, 0, len(stmts))
	for _, stmt := range stmts {
		if strings.TrimSpace(stmt) == "" {
			continue
		}
		el := newErrorListener()
		l := newQueriesParserListener()
		stmt := strings.TrimSpace(stmt)
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
		queries = append(queries, &Query{
			Stmt:    stmt,
			Selects: l.selects,
			Params:  l.params,
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
