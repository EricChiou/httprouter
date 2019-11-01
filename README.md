# httprouter
<pre><code>import "github.com/EricChiou/httprouter"</code></pre>
## Use fasthttp as http server
https://github.com/valyala/fasthttp

## Support methods
GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS

## Set Headers
Set headers by using
<pre><code>httprouter.SetHeader("Access-Control-Allow-Origin", "*")</code></pre>

## Url Path Rules
Only accept 0-9, a-z, A-Z  
Should start with "/"  
Should not end with "/"

## How to use
<pre><code>package main

import (
	"fmt"

	"github.com/EricChiou/httprouter"
)

func main() {
	httprouter.Get("/", handler)
	httprouter.Get("/demo", handler)
	httprouter.Get("/demo/id/demo2", handler)
	httprouter.Get("/demo/demo/demo/demo", handler)

	// path parameter
	httprouter.Get("/demo/:id/demo", handlerParam)
	httprouter.Get("/:id/demo", handlerParam)
	httprouter.Get("/demo/demo/demo2/:id", handlerParam)

	// duplicate path
	// httprouter.Get("/demo/demo", handler)

	// invalid character, only accept 0-9, a-z, A-Z
	// httprouter.Get("/demo/&", handler)
	// httprouter.Get("/demo/:!", handler)

	// wrong format
	// httprouter.Get("demo/demo", handler) // should start with "/"
	// httprouter.Get("/demo/demo/", handler) // should not end with "/"
	// httprouter.Get("/demo//demo", handler)

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// start http server
	if err := httprouter.ListenAndServe(":6200"); err != nil {
		panic(err)
	}

	// start https server
	// if err := httprouter.ListenAndServeTLS(":6200", "cert file path...", "key file path..."); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}

func handler(context *httprouter.Context) {
	fmt.Fprintf(context.Ctx, "url path: %s", string(context.Ctx.Path()))
}

func handlerParam(context *httprouter.Context) {
	id, _ := context.GetPathParam("id")
	fmt.Fprintf(context.Ctx, "url path: %s\nid: %s", string(context.Ctx.Path()), id)
}</code></pre>
