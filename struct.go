package httprouter

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

// Context is use to pass variables between middleware
type Context struct {
	Rep    http.ResponseWriter
	Req    *http.Request
	Ctx    *fasthttp.RequestCtx
	Params []param
}

// PathParam get path param
func (context Context) PathParam(key string) string {
	for _, param := range context.Params {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

type param struct {
	Key, Value string
}

type header struct {
	key   string
	value string
}

type nodeTrees struct {
	get, post, put, delete, patch, head, options *node
}

type node struct {
	path      string
	wildChild bool
	run       *func(*Context)
	children  []*node
}
