package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler/internal/antlrcql"
	"regexp"
	"sync"
)

type astQueryAnalyzer struct {
	mu            sync.Mutex
	queryAnalysis *astQueryAnalysis
}

func newQueryAnalyzer() *astQueryAnalyzer {
	return &astQueryAnalyzer{}
}

type astQueryAnalysis struct {
	selects  []string
	params   []string
	table    string
	keyspace string
}

func (qa *astQueryAnalyzer) analyze(ctx antlrcql.ICqlContext) *astQueryAnalysis {
	qa.mu.Lock()
	defer qa.mu.Unlock()
	qa.queryAnalysis = &astQueryAnalysis{}
	var t antlr.Tree = ctx
	t = t.GetChild(0)
	switch t := t.(type) {
	case *antlrcql.Select_Context:
		qa.visitSelect(t)
	case *antlrcql.InsertContext:
		qa.visitInsert(t)
	case *antlrcql.Delete_Context:
		qa.visitDelete(t)
	}
	q := qa.queryAnalysis
	qa.queryAnalysis = nil
	return q
}

func (qa *astQueryAnalyzer) visitInsert(ctx *antlrcql.InsertContext) {
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
	qa.queryAnalysis.params = params
	qa.queryAnalysis.table = table.GetText()
}

func (qa *astQueryAnalyzer) visitSelect(c *antlrcql.Select_Context) {
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
	qa.queryAnalysis.selects = selects
	qa.visitWhereSpec(whereSpec)
	qa.visitFromSpec(fromSpec)
}

func (qa *astQueryAnalyzer) visitDelete(c *antlrcql.Delete_Context) {
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
	qa.visitFromSpec(fromSpec)
	qa.visitWhereSpec(whereSpec)
}

func (qa *astQueryAnalyzer) visitFromSpec(c *antlrcql.FromSpecContext) {
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
		qa.queryAnalysis.table = tn.GetText()
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
		qa.queryAnalysis.table = tn.GetText()
		qa.queryAnalysis.keyspace = kn.GetText()
	}
}

func (qa *astQueryAnalyzer) visitWhereSpec(c *antlrcql.WhereSpecContext) {
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
	qa.queryAnalysis.params = params
}
