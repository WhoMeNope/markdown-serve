package main

import (
	"github.com/valyala/fasthttp"
)

func corecssHandler(ctx *fasthttp.RequestCtx, css []byte) {
	ctx.SetContentType("text/css")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(css)
}
