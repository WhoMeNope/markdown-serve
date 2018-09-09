package main

import (
	"github.com/gomarkdown/markdown/html"
	"github.com/valyala/fasthttp"
)

type myHandler struct {
	renderer *html.Renderer
	styles   []byte
}

func (h *myHandler) handleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/core.css":
		corecssHandler(ctx, h.styles)
	case "/story":
		storyHandler(ctx, h.renderer)
	case "/":
		rootHandler(ctx)
	default:
		ctx.Error("404 - Not Found", fasthttp.StatusNotFound)
	}
}
