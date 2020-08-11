package httprouter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

// FasthttpHandler fasthttp http handler
func FasthttpHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		for _, header := range headers {
			ctx.Response.Header.Set(header.key, header.value)
		}
		fasthttpMethodHandler(ctx)
	}
}

func fasthttpMethodHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case http.MethodGet:
		fasthttpPathHandler(ctx, trees.Get)
	case http.MethodPost:
		fasthttpPathHandler(ctx, trees.Post)
	case http.MethodPut:
		fasthttpPathHandler(ctx, trees.Put)
	case http.MethodDelete:
		fasthttpPathHandler(ctx, trees.Delete)
	case http.MethodPatch:
		fasthttpPathHandler(ctx, trees.Patch)
	case http.MethodHead:
		fasthttpPathHandler(ctx, trees.Head)
	case http.MethodOptions:
		fasthttpPathHandler(ctx, trees.Options)
	default:
		fmt.Fprintf(ctx, "404 page not found")
	}
}

func fasthttpPathHandler(ctx *fasthttp.RequestCtx, tree *node) {
	params := Params{}
	path := strings.SplitN(string(ctx.RequestURI()), "?", 2)[0]

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Ctx: ctx, Params: params})
		return
	}

	fmt.Fprintf(ctx, "404 page not found")
}
