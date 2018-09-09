package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
	fmt.Fprintln(ctx, "<body><ol>")

	for _, f := range files {
		fmt.Fprintln(ctx, "<li><a href='/story?name="+f.Name()+"'>"+f.Name()+"</a></li>")
	}

	fmt.Fprintln(ctx, "</ol></body>")

	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
