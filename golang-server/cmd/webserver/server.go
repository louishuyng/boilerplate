package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rz-server/internal/common/interfaces"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	util   *interfaces.Util
	server *http.Server
}

var _ interfaces.Server = (*Server)(nil)

func NewServer(util *interfaces.Util) *Server {
	router := mux.NewRouter()

	server := &http.Server{
		Addr:         ":" + PORT,
		WriteTimeout: WRITE_TIMEOUT,
		ReadTimeout:  READ_TIMEOUT,
		IdleTimeout:  IDLE_TIMEOUT,
		Handler:      router, // Pass our instance of gorilla/mux in.
		ErrorLog:     util.Logger,
	}

	return &Server{
		router: router,
		server: server,
		util:   util,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) GET(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, handler).Methods("GET")
}

func (s *Server) POST(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, handler).Methods("POST")
}

func (s *Server) PUT(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, handler).Methods("PUT")
}

func (s *Server) DELETE(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, handler).Methods("DELETE")
}

func (s *Server) RegisterMiddlewares(handlers []func(http.Handler) http.Handler) {
	s.util.Log.Info("Registering middlewares", map[string]any{
		"count": len(handlers),
	})

	for _, handler := range handlers {
		s.router.Use(handler)
	}
}

func (s *Server) Start() {
	s.util.Log.Info("Server started", map[string]any{
		"address": s.server.Addr,
		"port":    PORT,
	})

	if err := s.server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *Server) WaitForShutdown() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", GRACEFUL_TIMEOUT, "the duration for graceful shutdown")
	flag.Parse()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
