package httprouter

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func methodHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case fasthttp.MethodGet:
		pathHandler(ctx, trees.Get)
	case fasthttp.MethodPost:
		pathHandler(ctx, trees.Post)
	case fasthttp.MethodPut:
		pathHandler(ctx, trees.Put)
	case fasthttp.MethodDelete:
		pathHandler(ctx, trees.Delete)
	case fasthttp.MethodPatch:
		pathHandler(ctx, trees.Patch)
	case fasthttp.MethodHead:
		pathHandler(ctx, trees.Head)
	case fasthttp.MethodOptions:
		pathHandler(ctx, trees.Options)
	default:
		fmt.Fprintf(ctx, "404 page not found")
	}
}

func pathHandler(ctx *fasthttp.RequestCtx, tree *node) {
	params := Params{}
	path := string(ctx.RequestURI())

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Ctx: ctx, Params: params})
		return
	}

	fmt.Fprintf(ctx, "404 page not found")
}

func mapping(tree *node, path, pathSeg string, params *Params) *func(*Context) {
	if tree.wildChild {
		*params = append(Params{{Key: tree.path, Value: path}}, *params...)
	}

	if len(pathSeg) == 0 {
		if tree.run != nil {
			return tree.run
		}
		return nil
	}

	path, pathSeg = filterPath(pathSeg)
	for _, child := range tree.children {
		if path == child.path || child.wildChild {
			if run := mapping(child, path, pathSeg, params); run != nil {
				return run
			}
		}
	}

	return nil
}
