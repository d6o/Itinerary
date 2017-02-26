package itinerary

import (
	"net/http"
)

// IRequestMatcher tests if a request match with the condition.
type IRequestMatcher interface {
	Match(*http.Request) bool
}

// IPath represent one of the many paths a request can get.
type IPath interface {
	SetHandler(func(http.ResponseWriter, *http.Request)) IPath
	Match(*http.Request) bool
	AddMatcher(matcher IRequestMatcher) IPath
	Handler() http.Handler
}

// Path implements IPath with a MatcherList who will dictate if the request can go though that path.
type Path struct {
	handler     http.Handler
	matcherList []IRequestMatcher
}

// NewPath is the Path constructor.
func NewPath() *Path {
	return &Path{}
}

// Handler is the getter to the handler var which is the goal of the path.
func (r *Path) Handler() http.Handler {
	return r.handler
}

// SetHandler determine the goal of the request.
func (r *Path) SetHandler(f func(http.ResponseWriter, *http.Request)) IPath {
	r.handler = http.HandlerFunc(f)
	return r
}

// AddMatcher put a new condition to request take the path.
func (r *Path) AddMatcher(matcher IRequestMatcher) IPath {
	r.matcherList = append(r.matcherList, matcher)
	return r
}

// Match verify if the request can take this path.
func (r *Path) Match(req *http.Request) bool {
	for _, k := range r.matcherList {
		if k.Match(req) != true {
			return false
		}
	}
	return true
}
