package gee

import (
	"log"
	"net/http"
)

type Handler func(context *Context)


type RouterGroup struct {
	prefix string
	middlewares []Handler
	parent *RouterGroup
	engine *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup  {
	engine := group.engine

	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler Handler) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler Handler) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler Handler) {
	group.addRoute("POST", pattern, handler)
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter() }
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = [] *RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) GET(pattern string, handler Handler)  {
	e.router.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler Handler)  {
	e.router.addRouter("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	context := NewContext(writer, request)
	e.router.handle(context)
}

func (e *Engine) Run(addr string)  {
	http.ListenAndServe(addr, e)
}

