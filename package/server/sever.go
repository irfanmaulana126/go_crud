package server

import (
	"belajar/package/config"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	http   *http.Server
	Router *mux.Router
}

func NewServer(cfg *config.Config) *Server {
	r := mux.NewRouter()

	return &Server{
		http: &http.Server{
			Addr: cfg.ServerAddress,
		},
		Router: r,
	}
}

func (s *Server) RegisterRouter(handler http.Handler) {
	s.http.Handler = handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowCredentials())(handler)
}

func (s *Server) ListenAndServe() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	errc := make(chan error)
	go func() {
		log.Printf("HTTP Server listen on %s\n", s.http.Addr)
		errc <- s.http.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case err := <-errc:
		log.Printf("Error when listen on %s\n", s.http.Addr)
		return err
	case <-quit:
		log.Println("Shutting down the server")
		return s.http.Shutdown(ctx)
	}
}
