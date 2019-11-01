package httprouter

import (
	"github.com/valyala/fasthttp"
)

var headers []header
var trees = Trees{
	Get:     &node{path: "", wildChild: false, run: nil, children: []*node{}},
	Post:    &node{path: "", wildChild: false, run: nil, children: []*node{}},
	Put:     &node{path: "", wildChild: false, run: nil, children: []*node{}},
	Delete:  &node{path: "", wildChild: false, run: nil, children: []*node{}},
	Patch:   &node{path: "", wildChild: false, run: nil, children: []*node{}},
	Head:    &node{path: "", wildChild: false, run: nil, children: []*node{}},
	Options: &node{path: "", wildChild: false, run: nil, children: []*node{}},
}

// Get is use to build new get api
func Get(path string, run func(*Context)) {
	addRoute(fasthttp.MethodGet, trees.Get, path, run)
}

// Post is use to build new get api
func Post(path string, run func(*Context)) {
	addRoute(fasthttp.MethodPost, trees.Post, path, run)
}

// Put is use to build new get api
func Put(path string, run func(*Context)) {
	addRoute(fasthttp.MethodPut, trees.Put, path, run)
}

// Delete is use to build new get api
func Delete(path string, run func(*Context)) {
	addRoute(fasthttp.MethodDelete, trees.Delete, path, run)
}

// Patch is use to build new get api
func Patch(path string, run func(*Context)) {
	addRoute(fasthttp.MethodPatch, trees.Patch, path, run)
}

// Head is use to build new get api
func Head(path string, run func(*Context)) {
	addRoute(fasthttp.MethodHead, trees.Head, path, run)
}

// Options is use to build new get api
func Options(path string, run func(*Context)) {
	addRoute(fasthttp.MethodOptions, trees.Options, path, run)
}

// SetHeader add api response header
func SetHeader(key string, value string) {
	headers = append(headers, header{key: key, value: value})
}

// ListenAndServe start http server
func ListenAndServe(addr string) error {
	return fasthttp.ListenAndServe(addr, func(ctx *fasthttp.RequestCtx) {
		for _, header := range headers {
			ctx.Response.Header.Set(header.key, header.value)
		}
		methodHandler(ctx)
	})
}

// ListenAndServeTLS start https server
func ListenAndServeTLS(addr, certFile, keyFile string) error {
	return fasthttp.ListenAndServeTLS(addr, certFile, keyFile, func(ctx *fasthttp.RequestCtx) {
		for _, header := range headers {
			ctx.Response.Header.Set(header.key, header.value)
		}
		methodHandler(ctx)
	})
}
