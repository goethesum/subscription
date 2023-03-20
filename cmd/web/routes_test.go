package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()

	// verify existence
	chiRoutes := testRoutes.(chi.Router)

	for _, route := range routes {
		routesExists(t, chiRoutes, route)
	}
}

func routesExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes,
		func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			if route == foundRoute {
				found = true
			}
			return nil
		})

	if !found {
		t.Errorf("did not find %s in registred routes", route)
	}
}
