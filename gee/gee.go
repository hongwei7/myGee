package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %q", req.URL.Path)
	}
}

func (engine *Engine) Run() {
	http.ListenAndServe(":9999", engine)
}
