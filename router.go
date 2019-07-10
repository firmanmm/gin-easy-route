package gineasyroute

import (
	"github.com/gin-gonic/gin"
)

//Router base struct for generating complex route easily
type Router struct {
	routes      map[string]*route
	middlewares []*middleware
	engine      *gin.Engine
}

//AddRoute to be build, route require method that it listen to
//method here is the accepted HTTP method (ex : GET, POST, PUT, PATCH, etc)
//routes are connected after the middleware is built
func (r *Router) AddRoute(method, url string, handler gin.HandlerFunc) {
	var routeInstance *route
	if routeCache, ok := r.routes[url]; ok {
		routeInstance = routeCache
	} else {
		routeInstance = new(route)
		routeInstance.handler = make(map[string]gin.HandlerFunc)
		r.routes[url] = routeInstance
	}
	routeInstance.url = url
	routeInstance.handler[method] = handler
}

//AddMiddleware to be build, middleware are processed in FIFO (First In First Out) style
//so add order matters
func (r *Router) AddMiddleware(url string, handler gin.HandlerFunc) {
	middleware := new(middleware)
	middleware.url = url
	middleware.handler = handler
	r.middlewares = append(r.middlewares, middleware)
}

//Build the route and register it to injected Gin Engine Framework
func (r *Router) Build() {
	routeGraph := newGraph()
	routeGraph.Parse(r.routes, r.middlewares)
	r.processGraph(&r.engine.RouterGroup, routeGraph)
}

func (r *Router) processGraph(parent *gin.RouterGroup, routeGraph *graph) {
	group := parent.Group(routeGraph.path)
	for _, middleware := range routeGraph.middlewares {
		group.Use(middleware.handler)
	}
	if routeGraph.endpoint != nil {
		for method, handler := range routeGraph.endpoint.handler {
			group.Handle(method, "", handler)
		}
	}
	for _, child := range routeGraph.childs {
		r.processGraph(group, child)
	}
}

//NewRouter construct a new router builder with injected engine
func NewRouter(engine *gin.Engine) *Router {
	instance := new(Router)
	instance.engine = engine
	instance.routes = make(map[string]*route)
	instance.middlewares = make([]*middleware, 0)
	return instance
}
