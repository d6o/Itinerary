package itinerary

import (
	"net/http"
	"strings"
)

const (
	prefix = "itinerary"
)

// Router is a HTTP Server Handler implementation.
type Router struct {
	routes []IPath
	Prefix string
}

// NewRouter is the constructor to Router.
func NewRouter() *Router {
	return &Router{Prefix: prefix}
}

// NewPath creates a new alternative path that the request can take.
func (r *Router) NewPath(f func(http.ResponseWriter, *http.Request)) IPath {
	route := NewPath()
	r.routes = append(r.routes, route)
	return route.SetHandler(f)
}

// ServeHTTP find the right path and calls it handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := r.Match(req)
	if handler == nil {
		handler = http.NotFoundHandler()
	}
	handler.ServeHTTP(w, req)
}

// AddMatcher add a new condition in all paths already registered.
func (r *Router) AddMatcher(matcher IRequestMatcher) {
	for _, route := range r.routes {
		route.AddMatcher(matcher)
	}
}

// Match loop all paths until find the right one.
func (r *Router) Match(req *http.Request) http.Handler {
	for _, route := range r.routes {
		if !route.Match(req) {
			continue
		}
		return route.Handler()
	}
	return nil
}

// RouteToQuery makes the routes parts accessible trough Query String
func RouteToQuery(req *http.Request) {
	replace := "&itinerary="
	req.URL.RawQuery = req.URL.RawQuery + strings.Replace(strings.TrimRight(req.URL.Path, "/"), "/", replace, -1)
}
