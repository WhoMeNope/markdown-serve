package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/valyala/fasthttp"
)

func storyHandler(ctx *fasthttp.RequestCtx, renderer *html.Renderer) {
	args := ctx.QueryArgs()
	log.Println("Parsing query args : ", args)

	if args.Has("name") {
		name := string(args.Peek("name"))

		file, err := ioutil.ReadFile(os.Getenv("PATH_TO_STORIES") + "/" + name)
		if err != nil {
			log.Println("could not get file : ", name)
			ctx.Error("404 - Not Found", fasthttp.StatusNotFound)
		}

		html := markdown.ToHTML(file, nil, renderer)

		ctx.SetContentType("text/html")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(html)
	} else {
		ctx.Error("404 - Not Found", fasthttp.StatusNotFound)
	}
}
