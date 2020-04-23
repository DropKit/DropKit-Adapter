package routes

import (
	"net/http"

	auth "github.com/DropKit/DropKit-Adapter/controller/auth"
	db "github.com/DropKit/DropKit-Adapter/controller/db"

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
	register("POST", "/api/db/create", db.SQLCreate, nil)
	register("POST", "/api/db/insert", db.SQLInsert, nil)
	register("POST", "/api/db/select", db.SQLSelect, nil)
	register("POST", "/api/db/update", db.SQLUpdate, nil)
	register("POST", "/api/db/delete", db.SQLDelete, nil)

	register("POST", "/api/auth/grant", auth.AuthGrant, nil)
	register("POST", "/api/auth/revoke", auth.AuthRevoke, nil)
	register("POST", "/api/auth/verify", auth.AuthVerify, nil)
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
