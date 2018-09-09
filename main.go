package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gomarkdown/markdown/html"
	"github.com/valyala/fasthttp"
)

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
		Title: "story tale",
		CSS:   "/" + os.Getenv("STYLEFILE"),
		Flags: htmlFlags,
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
