package cql

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cql/parser"
)

type listener struct {
	*parser.BaseCqlParserListener
}

// func (l *listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText()) // Print every rule entered
// }

func (l *listener) EnterSelectElement(ctx *parser.SelectElementContext) {
	fmt.Println(ctx.GetText())
}

func Parse(cql string) error {
	lexer := parser.NewCqlLexer(antlr.NewInputStream(cql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCqlParser(stream)
	p.RemoveErrorListeners()
	l := &listener{}
	errl := newErrorListener()
	p.AddErrorListener(errl)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Cql())
	return errors.Join(errl.errors...)
}

type errorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorMessage := fmt.Errorf("line %d:%d %s", line, column, msg)
	l.errors = append(l.errors, errorMessage)
}

func newErrorListener() *errorListener {
	return &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}
