package masterRoutes

import (
	"belajar/package/manager"

	masterRoute "belajar/api/master/delivery/route"

	"github.com/gorilla/mux"
)

func NewRoutes(r *mux.Router, mgr manager.Manager) {
	apiV1Master := r.PathPrefix("/v1").Subrouter()

	masterRoute.NewAccessRoute(mgr, apiV1Master)
	masterRoute.NewRolesRoute(mgr, apiV1Master)
	masterRoute.NewPermissionsRoute(mgr, apiV1Master)

}
