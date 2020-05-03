package routes

import (
	"net/http"

	db "github.com/DropKit/DropKit-Adapter/controller/db"
	health "github.com/DropKit/DropKit-Adapter/controller/health"
	payment "github.com/DropKit/DropKit-Adapter/controller/payment"
	permission "github.com/DropKit/DropKit-Adapter/controller/permission"
	user "github.com/DropKit/DropKit-Adapter/controller/user"

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
	register("GET", "/health/ping", health.Ping, nil)
	register("GET", "/health/dependency", health.DependencyCheck, nil)

	register("GET", "/user/create", user.CreateUser, nil)

	register("POST", "/db/create", db.SQLCreate, nil)
	register("POST", "/db/insert", db.SQLInsert, nil)
	register("POST", "/db/select", db.SQLSelect, nil)
	register("POST", "/db/update", db.SQLUpdate, nil)
	register("POST", "/db/delete", db.SQLDelete, nil)

	register("POST", "/permission/grant/admin", permission.GrantAdmin, nil)
	register("POST", "/permission/grant/maintainer", permission.GrantMaintainer, nil)
	register("POST", "/permission/grant/user", permission.GrantUser, nil)

	register("POST", "/permission/revoke/admin", permission.RevokeAdmin, nil)
	register("POST", "/permission/revoke/maintainer", permission.RevokeMaintainer, nil)
	register("POST", "/permission/revoke/user", permission.RevokeUser, nil)

	register("POST", "/permission/verify/admin", permission.VerifyAdmin, nil)
	register("POST", "/permission/verify/maintainer", permission.VerifyMaintainer, nil)
	register("POST", "/permission/verify/user", permission.VerifyUser, nil)

	register("POST", "/payment/mint", payment.MintToken, nil)
	register("POST", "/payment/burn", payment.BurnToken, nil)
	register("POST", "/payment/transfer", payment.TransferToken, nil)
	register("POST", "/payment/balance", payment.GetBalance, nil)
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
