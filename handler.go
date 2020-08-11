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
		pathHandler(rep, req, trees.Get)
	case http.MethodPost:
		pathHandler(rep, req, trees.Post)
	case http.MethodPut:
		pathHandler(rep, req, trees.Put)
	case http.MethodDelete:
		pathHandler(rep, req, trees.Delete)
	case http.MethodPatch:
		pathHandler(rep, req, trees.Patch)
	case http.MethodHead:
		pathHandler(rep, req, trees.Head)
	case http.MethodOptions:
		pathHandler(rep, req, trees.Options)
	default:
		fmt.Fprintf(rep, "404 page not found")
	}
}

func pathHandler(rep http.ResponseWriter, req *http.Request, tree *node) {
	params := Params{}
	path := req.RequestURI

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Rep: rep, Req: req, Params: params})
		return
	}

	fmt.Fprintf(rep, "404 page not found")
}
