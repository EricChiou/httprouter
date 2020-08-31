package httprouter

import (
	"fmt"
	"net/http"
)

// HTTPHandler net/http http handler
func HTTPHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rep http.ResponseWriter, req *http.Request) {
		for _, header := range headers {
			rep.Header().Set(header.key, header.value)
		}
		methodHandler(rep, req)
	})
	return mux
}

func methodHandler(rep http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		pathHandler(rep, req, trees.get)
	case http.MethodPost:
		pathHandler(rep, req, trees.post)
	case http.MethodPut:
		pathHandler(rep, req, trees.put)
	case http.MethodDelete:
		pathHandler(rep, req, trees.delete)
	case http.MethodPatch:
		pathHandler(rep, req, trees.patch)
	case http.MethodHead:
		pathHandler(rep, req, trees.head)
	case http.MethodOptions:
		pathHandler(rep, req, trees.options)
	default:
		fmt.Fprintf(rep, "404 page not found")
	}
}

func pathHandler(rep http.ResponseWriter, req *http.Request, tree *node) {
	params := []param{}
	path := req.RequestURI

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Rep: rep, Req: req, Params: params})
		return
	}

	fmt.Fprintf(rep, "404 page not found")
}
