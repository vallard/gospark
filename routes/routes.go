package routes

import (
	"net/http"

	"github.com/vallard/gospark/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GitCommit",
		"POST",
		"/v1/commit",
		handlers.GitCommit,
	},
	Route{
		"GitHello",
		"GET",
		"/v1/hello",
		handlers.GitHello,
	},
}
