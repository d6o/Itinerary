package itinerary

import (
	"net/http"
)

type IRequestMatcher interface {
	Match(*http.Request) bool
}

type IRoute interface {
	SetHandler(func(http.ResponseWriter, *http.Request)) IRoute
	Match(*http.Request) bool
	AddMatcher(matcher IRequestMatcher) IRoute
	Handler() http.Handler
}

type Route struct {
	handler     http.Handler
	matcherList []IRequestMatcher
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) Handler() http.Handler {
	return r.handler
}

func (r *Route) SetHandler(f func(http.ResponseWriter, *http.Request)) IRoute {
	r.handler = http.HandlerFunc(f)
	return r
}

func (r *Route) AddMatcher(matcher IRequestMatcher) IRoute {
	r.matcherList = append(r.matcherList, matcher)
	return r
}

func (r *Route) Match(req *http.Request) bool {
	for _, k := range r.matcherList {
		if k.Match(req) != true {
			return false
		}
	}
	return true
}
