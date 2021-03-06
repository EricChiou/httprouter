package httprouter

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func checkFormat(method, path string) error {
	// check path is valid
	if len(path) == 0 {
		msg := "path error, path can not be empty."
		log.Println(msg)
		return errors.New(msg)
	}

	if path[0:1] != "/" {
		msg := "path error, " + method + ": '" + path + "', path must begin with '/'."
		log.Println(msg)
		return errors.New(msg)
	}

	if len(path) == 1 {
		return nil
	}

	paths := strings.Split(path[1:], "/")
	for _, p := range paths {
		if len(p) > 0 && p[0:1] == ":" {
			p = p[1:]
		}
		if len(p) == 0 {
			msg := "path error, " + method + ": '" + path + "', path has wrong format."
			log.Println(msg)
			return errors.New(msg)
		}
		match, _ := regexp.MatchString("^[0-9a-zA-Z-_]+$", p)
		if !match {
			msg := "path error, " + method + ": '" + path + "', path has invalid character, only accept 0-9, a-z, A-Z, -, _."
			log.Println(msg)
			return errors.New(msg)
		}
	}
	return nil
}

func checkDuplicate(tree *node, path, pathSeg string) bool {
	if len(pathSeg) == 0 {
		if tree.run != nil {
			return true
		}
	} else {
		path, pathSeg = filterPath(pathSeg)
		wildChild := (path[0:1] == ":")
		if wildChild {
			path = path[1:]
		}

		for _, child := range tree.children {
			if path == child.path || wildChild || child.wildChild {
				if checkDuplicate(child, path, pathSeg) {
					return true
				}
			}
		}
	}

	return false
}

func addPath(tree *node, path, pathSeg string, run func(*Context)) bool {
	if len(pathSeg) == 0 {
		tree.run = &run
		return true
	}

	path, pathSeg = filterPath(pathSeg)
	wildChild := (path[0:1] == ":")
	if wildChild {
		path = path[1:]
	}

	for _, child := range tree.children {
		if path == child.path && wildChild == child.wildChild {
			if addPath(child, path, pathSeg, run) {
				return true
			}
		}
	}

	newChild := node{path: path, wildChild: wildChild, run: nil, children: []*node{}}
	tree.children = append(tree.children, &newChild)
	if addPath(&newChild, path, pathSeg, run) {
		return true
	}

	return false
}

func filterPath(path string) (string, string) {
	for i, c := range path {
		if c == 47 {
			return path[:i], path[i+1:]
		}
	}
	return path, ""
}

func mapping(tree *node, path, pathSeg string, params *[]param) *func(*Context) {
	if tree.wildChild {
		*params = append([]param{{Key: tree.path, Value: path}}, *params...)
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
