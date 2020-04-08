package routes

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/controller"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("POST", "/api/db/create", controller.SQLCreate, nil)
	register("POST", "/api/db/insert", controller.SQLInsert, nil)
	register("POST", "/api/db/select", controller.SQLSelect, nil)

	register("POST", "/api/auth/grant", controller.AuthGrant, nil)
	register("POST", "/api/auth/revoke", controller.AuthRevoke, nil)
	register("POST", "/api/auth/verify", controller.AuthVerify, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
