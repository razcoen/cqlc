package cqlc

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/internal/antlrcql"
	"reflect"
)

type baseCQLParserListener struct {
	*antlrcql.BaseCQLParserListener
	debug bool
}

func (l *baseCQLParserListener) EnterEveryRule(c antlr.ParserRuleContext) {
	if !l.debug {
		return
	}
	fmt.Println(reflect.TypeOf(c).String(), c.GetText())
}

func (l *baseCQLParserListener) VisitTerminal(node antlr.TerminalNode) {
	if !l.debug {
		return
	}
	fmt.Println(reflect.TypeOf(node).String(), node.GetText())
}
