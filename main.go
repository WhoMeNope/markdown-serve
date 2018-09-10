package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/valyala/fasthttp"
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

func main() {
	//load styles
	styles, err := ioutil.ReadFile(os.Getenv("PATH_TO_STYLES") + "/" + os.Getenv("STYLEFILE"))
	if err != nil {
		log.Fatal("Error reading " + os.Getenv("STYLEFILE"))
	}
	log.Println("loaded stylesheets")

	//setup renderer
	htmlFlags := html.CommonFlags | html.CompletePage
	opts := html.RendererOptions{
		Title:          "story tale",
		CSS:            "/" + os.Getenv("STYLEFILE"),
		Flags:          htmlFlags,
		RenderNodeHook: customRenderHook,
	}
	renderer := html.NewRenderer(opts)

	//serve
	myHandler := myHandler{
		renderer: renderer,
		styles:   styles,
	}

	log.Println("serving at :" + os.Getenv("PORT"))
	fasthttp.ListenAndServe(":"+os.Getenv("PORT"), myHandler.handleFastHTTP)
}
