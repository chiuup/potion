package potion

import "net/http"

// Route ...
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an alias of Route array.
type Routes []Route
