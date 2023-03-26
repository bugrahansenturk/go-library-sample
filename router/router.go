package router

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) RegisterRoute(path string, handler http.HandlerFunc) {
	route := Route{
		Path:    path,
		Handler: handler,
	}

	r.routes = append(r.routes, route)
}

func (r *Router) RegisterRoutes(routes []Route) {
	for _, route := range routes {
		r.RegisterRoute(route.Path, route.Handler)
	}
}

func (r *Router) SetupRoutes(mux *http.ServeMux) {
	for _, route := range r.routes {
		mux.HandleFunc(route.Path, route.Handler)
	}
}
