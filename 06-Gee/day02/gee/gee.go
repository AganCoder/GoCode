package gee


import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.router.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.addRouter("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}