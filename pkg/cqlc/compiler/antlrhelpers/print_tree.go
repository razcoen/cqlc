package antlrhelpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/xlab/treeprint"
)

func PrintTree(t antlr.Tree) {
	p := treeprint.New()
	createTreePrint(p, t)
	fmt.Println(p.String())
}

func createTreePrint(p treeprint.Tree, t antlr.Tree) {
	type GetTexter interface {
		GetText() string
	}
	getTexter, ok := t.(GetTexter)
	if !ok {
		return
	}
	p1 := p.AddBranch(strings.Join([]string{reflect.TypeOf(t).String(), " :: ", getTexter.GetText()}, ""))
	for _, c := range t.GetChildren() {
		createTreePrint(p1, c)
	}
}
