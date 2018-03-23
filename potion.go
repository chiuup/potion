package potion

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// Potion ...
type Potion struct {
	router *httprouter.Router
}

// New creates a new Potion.
func New() *Potion {
	router := httprouter.New()
	router.MethodNotAllowed = loggerHandler(http.HandlerFunc(methodNotAllowedHandler))
	router.NotFound = loggerHandler(http.HandlerFunc(notFoundHandler))
	return &Potion{router}
}

// RegisterRoutes register a group of routes to a Potion.
// Logging and panic recovery handlers are added.
func (p *Potion) RegisterRoutes(routes *Routes) {
	for _, route := range *routes {
		p.router.Handler(route.Method, route.Pattern, loggerHandler(route.HandlerFunc))
	}
}

// ServeHTTP ...
func (p Potion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}
