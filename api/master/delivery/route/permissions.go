package masterRoute

import (
	masterHandler "belajar/api/master/delivery/handler"
	"belajar/package/manager"

	"github.com/gorilla/mux"
)

func NewPermissionsRoute(mgr manager.Manager, route *mux.Router) {
	permissionsHandler := masterHandler.NewPermissionsHandler(mgr)

	route.Handle("/permissions", permissionsHandler.GetPermissions()).Methods("GET")
}
