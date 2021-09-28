package gee

import (
	"net/http"
)

type Handler func(context *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
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

