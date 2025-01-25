package compiler

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler/internal/antlrcql"
)

func antlrParse(text string) ([]antlrcql.ICqlContext, error) {
	p, el := newAntlrCQLParser(text)
	cqls := p.Cqls().AllCql()
	if len(el.errors) > 0 {
		return nil, errors.Join(el.errors...)
	}
	if len(cqls) > 0 {
		return cqls, nil
	}
	// Try using the singular cql parsing in case the plural doesn't work.
	// TODO: Figure out why this happens and why does this fix the FROZEN type test for the schema parser.
	p, el = newAntlrCQLParser(text)
	cql := p.Cql()
	if len(el.errors) > 0 {
		return nil, errors.Join(el.errors...)
	}
	return []antlrcql.ICqlContext{cql}, nil
}

func newAntlrCQLParser(text string) (*antlrcql.CQLParser, *errorListener) {
	el := newErrorListener()
	lexer := antlrcql.NewCQLLexer(antlr.NewInputStream(text))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := antlrcql.NewCQLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	return p, el
}

type errorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func newErrorListener() *errorListener {
	return &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	token, ok := offendingSymbol.(antlr.Token)
	var tokenText string
	if ok {
		tokenText = token.GetText()
	}
	// TODO: The following are ignored but requires fix through the g4 syntax files.
	if tokenText == "<EOF>" || tokenText == "<" {
		return
	}
	err := fmt.Errorf(`syntax error "%s" in line %d:%d %s`, tokenText, line, column, msg)
	l.errors = append(l.errors, err)
}
