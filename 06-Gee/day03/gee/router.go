package gee

import (
	"fmt"
	"strings"
)

// Router 用于用户路由选择
// 内部使用类型-因此无须暴露出去
// Root for Method

type router struct {
	roots map[string]*node
	router map[string]Handler
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*node),
		router: make(map[string]Handler),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}


func (r *router) addRouter(method string, pattern string, handler Handler)  {

	parts := parsePattern(pattern)

	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.router[key] = handler
}

func (r *router) getRouter(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return  nil, nil
	}
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}


func (r *router) handle(context *Context)  {
	n, params := r.getRouter(context.Method, context.Path)

	if n != nil {
		context.Params = params
		key := context.Method + "-" + n.pattern
		fmt.Println(key)

		r.router[key](context)
	} else {
		fmt.Fprintf(context.Writer, "404 NOT Found")
	}
}