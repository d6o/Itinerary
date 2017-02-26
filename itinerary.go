package itinerary

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	prefix = "itinerary"
)

// Router is a HTTP Server Handler implementation.
type Router struct {
	routes []IRoute
	Prefix string
}

// NewRouter is the constructor to Router.
func NewRouter() *Router {
	return &Router{Prefix: prefix}
}

// HandleFunc Gorilla Mux compatible function to insert new routes to the router.
func (r *Router) HandleFunc(f func(http.ResponseWriter, *http.Request)) IRoute {
	route := NewRoute()
	r.routes = append(r.routes, route)
	return route.SetHandler(f)
}

// ServeHTTP find the right route and calls it handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.HandleRoute(req)
	handler := r.Match(req)

	if handler == nil {
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, req)
}

// Match loop all routes until find the right one.
func (r *Router) Match(req *http.Request) http.Handler {
	for _, route := range r.routes {
		if !route.Match(req) {
			continue
		}
		return route.Handler()
	}
	return nil
}

func (r *Router) HandleRoute(req *http.Request) {
	req.URL.Path = strings.Trim(req.URL.Path, "/")
	if req.URL.Path == "" {
		return
	}
	parts := strings.Split(req.URL.Path, "/")
	value := make(url.Values)
	for i, part := range parts {
		value.Set(r.Prefix+strconv.Itoa(i), part)
	}
	req.URL.RawQuery = url.Values(value).Encode() + "&" + req.URL.RawQuery
}
