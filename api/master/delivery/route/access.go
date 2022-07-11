package masterRoute

import (
	masterHandler "belajar/api/master/delivery/handler"
	"belajar/package/manager"

	"github.com/gorilla/mux"
)

func NewAccessRoute(mgr manager.Manager, route *mux.Router) {
	accessHandler := masterHandler.NewAccessHandler(mgr)

	route.Handle("/access", accessHandler.GetAccess()).Methods("GET")
}
