package masterRoute

import (
	masterHandler "belajar/api/master/delivery/handler"
	"belajar/package/manager"

	"github.com/gorilla/mux"
)

func NewRolesRoute(mgr manager.Manager, route *mux.Router) {
	rolesHandler := masterHandler.NewRolesHandler(mgr)

	route.Handle("/roles", rolesHandler.GetRoles()).Methods("GET")
}
