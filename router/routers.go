package routes

import (
	"net/http"

	db "github.com/DropKit/DropKit-Adapter/controller/db"
	health "github.com/DropKit/DropKit-Adapter/controller/health"
	payment "github.com/DropKit/DropKit-Adapter/controller/payment"
	permission "github.com/DropKit/DropKit-Adapter/controller/permission"
	role "github.com/DropKit/DropKit-Adapter/controller/role"
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
	// Todo: Merge two service into one, just check the whole services.
	register("GET", "/health/ping", health.PerformHealthCheck, nil)
	register("GET", "/health/dependency", health.CheckDependencyServices, nil)

	register("GET", "/user/create", user.GenerateRandomAccount, nil)

	register("POST", "/db/create", db.HandleDBCreation, nil)
	register("POST", "/db/insert", db.HandleDBInsertion, nil)
	register("POST", "/db/select", db.HandleDBSelection, nil)
	register("POST", "/db/update", db.HandleDBUpdate, nil)
	register("POST", "/db/delete", db.HandleDBDeletion, nil)

	register("POST", "/permission/grant/table/owner", permission.GrantTableOwner, nil)
	register("POST", "/permission/grant/table/maintainer", permission.GrantTableMaintainer, nil)
	register("POST", "/permission/grant/table/viewer", permission.GrantTableViewer, nil)

	register("POST", "/permission/revoke/table/owner", permission.RevokeTableOwner, nil)
	register("POST", "/permission/revoke/table/maintainer", permission.RevokeTableMaintainer, nil)
	register("POST", "/permission/revoke/table/viewer", permission.RevokeTableViewer, nil)

	register("POST", "/permission/verify/table/owner", permission.VerifyTableOwner, nil)
	register("POST", "/permission/verify/table/maintainer", permission.VerifyTableMaintainer, nil)
	register("POST", "/permission/verify/table/viewer", permission.TableViewer, nil)

	register("POST", "/payment/mint", payment.MintToken, nil)
	register("POST", "/payment/burn", payment.BurnToken, nil)
	register("POST", "/payment/transfer", payment.TransferToken, nil)
	register("POST", "/payment/balance", payment.GetAccountBalance, nil)

	register("POST", "/role/create", role.CreateColumnRole, nil)
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
