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
		fasthttpPathHandler(ctx, trees.get)
	case http.MethodPost:
		fasthttpPathHandler(ctx, trees.post)
	case http.MethodPut:
		fasthttpPathHandler(ctx, trees.put)
	case http.MethodDelete:
		fasthttpPathHandler(ctx, trees.delete)
	case http.MethodPatch:
		fasthttpPathHandler(ctx, trees.patch)
	case http.MethodHead:
		fasthttpPathHandler(ctx, trees.head)
	case http.MethodOptions:
		fasthttpPathHandler(ctx, trees.options)
	default:
		fmt.Fprintf(ctx, "404 page not found")
	}
}

func fasthttpPathHandler(ctx *fasthttp.RequestCtx, tree *node) {
	params := []param{}
	path := strings.SplitN(string(ctx.RequestURI()), "?", 2)[0]

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Ctx: ctx, Params: params})
		return
	}

	fmt.Fprintf(ctx, "404 page not found")
}
