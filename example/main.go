package main

import (
	"fmt"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	handler := func(context *httprouter.Context) {
		fmt.Fprintf(context.Ctx, "url path: %s", string(context.Ctx.Path()))
	}
	handlerParam := func(context *httprouter.Context) {
		id := context.PathParam("id")
		fmt.Fprintf(context.Ctx, "url path: %s\nid: %s", string(context.Ctx.Path()), id)
	}

	httprouter.Get("/", handler)
	httprouter.Get("/path", handler)
	httprouter.Get("/path/id/path2", handler)
	httprouter.Get("/path/path/path2", handler)

	// path parameter
	httprouter.Get("/path/:id/path", handlerParam)
	httprouter.Get("/:id/path", handlerParam)
	httprouter.Get("/path/path/path2/:id", handlerParam)

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// fasthttp http server
	if err := fasthttp.ListenAndServe(":6200", httprouter.FasthttpHandler()); err != nil {
		panic(err)
	}

	// fasthttp https server
	// if err := fasthttp.ListenAndServeTLS(":6200", "cert file path...", "key file path...", httprouter.FasthttpHandler()); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}
