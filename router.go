package httprouter

import (
	"errors"
	"log"
	"net/http"
)

var headers []header
var trees = nodeTrees{
	get:     &node{path: "", wildChild: false, run: nil, children: []*node{}},
	post:    &node{path: "", wildChild: false, run: nil, children: []*node{}},
	put:     &node{path: "", wildChild: false, run: nil, children: []*node{}},
	delete:  &node{path: "", wildChild: false, run: nil, children: []*node{}},
	patch:   &node{path: "", wildChild: false, run: nil, children: []*node{}},
	head:    &node{path: "", wildChild: false, run: nil, children: []*node{}},
	options: &node{path: "", wildChild: false, run: nil, children: []*node{}},
}

// Get is use to build new get api
func Get(path string, run func(*Context)) error {
	return addRoute(http.MethodGet, trees.get, path, run)
}

// Post is use to build new get api
func Post(path string, run func(*Context)) error {
	return addRoute(http.MethodPost, trees.post, path, run)
}

// Put is use to build new get api
func Put(path string, run func(*Context)) error {
	return addRoute(http.MethodPut, trees.put, path, run)
}

// Delete is use to build new get api
func Delete(path string, run func(*Context)) error {
	return addRoute(http.MethodDelete, trees.delete, path, run)
}

// Patch is use to build new get api
func Patch(path string, run func(*Context)) error {
	return addRoute(http.MethodPatch, trees.patch, path, run)
}

// Head is use to build new get api
func Head(path string, run func(*Context)) error {
	return addRoute(http.MethodHead, trees.head, path, run)
}

// Options is use to build new get api
func Options(path string, run func(*Context)) error {
	return addRoute(http.MethodOptions, trees.options, path, run)
}

// SetHeader add api response header
func SetHeader(key string, value string) {
	headers = append(headers, header{key: key, value: value})
}

func addRoute(method string, tree *node, path string, run func(*Context)) error {
	if err := checkFormat(method, path); err != nil {
		return err
	}

	if checkDuplicate(tree, "", path[1:]) {
		msg := "path duplicated, " + method + ": '" + path
		log.Println(msg)
		return errors.New(msg)
	}

	if !addPath(tree, "", path[1:], run) {
		msg := "add path fail, " + method + ": '" + path
		log.Fatalln(msg)
		return errors.New(msg)
	}

	return nil
}
