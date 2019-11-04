# httprouter
```go
import "github.com/EricChiou/httprouter"
```
## Support both net/http and fasthttp
net/http - https://golang.org/pkg/net/http/  
fasthttp - https://github.com/valyala/fasthttp

## Support methods
GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS

## Set Headers
Set headers by using
```go
httprouter.SetHeader("Access-Control-Allow-Origin", "*")
```

## Url Path Rules
- Only accept 0-9, a-z, A-Z  
- Should start with "/"  
- Should not end with "/"

## How to use
### net/http
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/EricChiou/httprouter"
)

func main() {
	hadler := func(context *httprouter.Context) {
		fmt.Fprintf(context.Rep, "url path: %s", context.Req.RequestURI)
	}
	handlerParam := func(context *httprouter.Context) {
		id, _ := context.GetPathParam("id")
		fmt.Fprintf(context.Rep, "url path: %s\nid: %s", context.Req.RequestURI, id)
	}
	
	httprouter.Get("/", handler)
	httprouter.Get("/path", handler)
	httprouter.Get("/path/id/path2", handler)
	httprouter.Get("/path/path/path2", handler)

	// path parameter
	httprouter.Get("/path/:id/path", handlerParam)
	httprouter.Get("/:id/path", handlerParam)
	httprouter.Get("/path/path/path2/:id", handlerParam)

	// duplicate path
	// httprouter.Get("/path/path", handler)

	// invalid character, only accept 0-9, a-z, A-Z
	// httprouter.Get("/path/&", handler)
	// httprouter.Get("/path/:!", handler)

	// wrong format
	// httprouter.Get("path/path", handler) // should start with "/"
	// httprouter.Get("/path/path/", handler) // should not end with "/"
	// httprouter.Get("/path//path", handler)

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// net/http http server
	if err := http.ListenAndServe(":6200", httprouter.HTTPHandler()); err != nil {
		panic(err)
	}

	// net/http https server
	// if err := http.ListenAndServeTLS(":6200", "cert file path...", "key file path...", httprouter.HTTPHandler()); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}
```
### fasthttp
```go
package main

import (
	"fmt"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	hadler := func(context *httprouter.Context) {
		fmt.Fprintf(context.Ctx, "url path: %s", string(context.Ctx.Path()))
	}
	handlerParam := func(context *httprouter.Context) {
		id, _ := context.GetPathParam("id")
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
```
