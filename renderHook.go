package main

import (
	"io"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

func customRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	// skip all nodes that are not text nodes
	if _, ok := node.(*ast.Text); !ok {
		return ast.GoToNext, false
	}

	textNode := node.(*ast.Text)

	// custom rendering logic for ast.Text
	for _, v := range strings.Split(string(textNode.Leaf.Literal), "\n") {
		w.Write([]byte(v + "</br>"))
	}

	// return (ast.GoToNext, true) to tell html renderer to skip rendering this node
	return ast.GoToNext, true
}
