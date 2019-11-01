# httprouter
## Use fasthttp as http server
https://github.com/valyala/fasthttp

## Support methods
GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS

## Set Headers
Set headers by use "httprouter.SetHeader("Access-Control-Allow-Origin", "*")"

## Url Path only accept 0-9, a-z, A-Z

## How to use
<pre><code>package main

import (
	"fmt"

	"github.com/EricChiou/httprouter"
)

func main() {
	httprouter.Get("/", demo)
	httprouter.Get("/demo", demo)
	httprouter.Get("/demo/id/demo2", demo)
	httprouter.Get("/demo/demo/demo/demo", demo)

	// path parameter
	httprouter.Get("/demo/:id/demo", demoParam)
	httprouter.Get("/:id/demo", demoParam)
	httprouter.Get("/demo/demo/demo2/:id", demoParam)

	// duplicate path
	// httprouter.Get("/demo/demo", demo)

	// invalid character, only accept 0-9, a-z, A-Z
	// httprouter.Get("/demo/&", demo)
	// httprouter.Get("/demo/:!", demo)

	// wrong format
	// httprouter.Get("demo/demo", demo) // should start with "/"
	// httprouter.Get("/demo/demo/", demo) // should not end with "/"
	// httprouter.Get("/demo//demo", demo)

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

func demo(context *httprouter.Context) {
	fmt.Fprintf(context.Ctx, "url path: %s", string(context.Ctx.Path()))
}

func demoParam(context *httprouter.Context) {
	id, _ := context.GetPathParam("id")
	fmt.Fprintf(context.Ctx, "url path: %s\nid: %s", string(context.Ctx.Path()), id)
}</code></pre>
