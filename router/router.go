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

func (router *Router) RegisterRoute(path string, handler http.HandlerFunc) {
	route := Route{
		Path:    path,
		Handler: handler,
	}

	router.routes = append(router.routes, route)
}

func (router *Router) RegisterRoutes(routes []Route) {
	for _, route := range routes {
		router.RegisterRoute(route.Path, route.Handler)
	}
}

func (router *Router) SetupRoutes(mux *http.ServeMux) {
	for _, route := range router.routes {
		mux.HandleFunc(route.Path, route.Handler)
	}
}
