package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

func rootHandler(ctx *fasthttp.RequestCtx) {
	files, err := ioutil.ReadDir(os.Getenv("PATH_TO_STORIES"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(ctx,
		`<head>
			<title>tales</title>
			<link rel="stylesheet" type="text/css" href="/core.css">
		</head>`)
	fmt.Fprintln(ctx, "<body><div class='list'><ul>")

	for _, f := range files {
		name := stripExtension(f.Name())
		fmt.Fprintln(ctx, "<li><a href='/story?name="+f.Name()+"'>"+name+"</a></li>")
	}

	fmt.Fprintln(ctx, "</ul></div></body>")

	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func stripExtension(s string) string {
	parts := strings.Split(s, ".")

	// no extension
	if len(parts) <= 1 {
		return s
	}

	// concat without extension
	name := ""
	for i := 0; i < len(parts)-1; i++ {
		if name != "" {
			name += "."
		}
		name += parts[i]
	}
	return name
}
