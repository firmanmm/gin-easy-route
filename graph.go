package gineasyroute

import (
	"strings"
)

type graph struct {
	path        string
	endpoint    *route
	middlewares []*middleware
	childs      map[string]*graph
}

func (g *graph) AddRoute(paths []string, data *route) {
	path := paths[0]
	if len(paths) < 2 {
		g.path = path
		g.endpoint = data
	} else {
		path = paths[1]
		node, ok := g.childs[path]
		if !ok {
			node = newGraph()
			node.path = path
			g.childs[path] = node
		}
		node.AddRoute(paths[1:], data)
	}
}

func (g *graph) AddMiddleware(paths []string, data *middleware) {
	path := paths[0]
	if len(paths) < 2 {
		g.middlewares = append(g.middlewares, data)
	} else {
		path = paths[1]
		node, ok := g.childs[path]
		if !ok {
			node = newGraph()
			node.path = path
			g.childs[path] = node
		}
		node.AddMiddleware(paths[1:], data)
	}
}

func (g *graph) Parse(routes map[string]*route, middlewares []*middleware) {
	for _, route := range routes {
		param := strings.Split(route.url, "/")
		if param[1] == "" {
			param = param[:1]
		}
		g.AddRoute(param, route)
	}
	for _, middleware := range middlewares {
		param := strings.Split(middleware.url, "/")
		if param[1] == "" {
			param = param[:1]
		}
		g.AddMiddleware(param, middleware)
	}
}

func newGraph() *graph {
	instance := new(graph)
	instance.middlewares = make([]*middleware, 0)
	instance.childs = make(map[string]*graph)
	return instance
}
