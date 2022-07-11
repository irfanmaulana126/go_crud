package main

import (
	"belajar/package/manager"
	"belajar/package/server"
	"fmt"
	"os"

	masterRoutes "belajar/api/master/delivery"
)

func run() error {
	mgr, err := manager.NewInit()
	if err != nil {
		return err
	}

	server := server.NewServer(mgr.GetConfig())
	server.Router.Use(mgr.GetMiddleware().InitLog)

	masterRoutes.NewRoutes(server.Router, mgr)

	server.RegisterRouter(server.Router)
	return server.ListenAndServe()
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
