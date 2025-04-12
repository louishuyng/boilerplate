package playground_setup

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rz-server/helpers"
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

const TEST = 15 * time.Second

func NewServer(util *interfaces.Util) *Server {
	router := mux.NewRouter()

	port := util.Environment.GetEnv("server", "PORT")
	writeTimeoutSeconds := util.Environment.GetEnv("server", "WRITE_TIMEOUT_SECONDS")
	readTimeoutSeconds := util.Environment.GetEnv("server", "READ_TIMEOUT_SECONDS")
	idleTimeoutSeconds := util.Environment.GetEnv("server", "IDLE_TIMEOUT_SECONDS")

	server := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Duration(helpers.StrToInt(writeTimeoutSeconds)) * time.Second,
		ReadTimeout:  time.Duration(helpers.StrToInt(readTimeoutSeconds)) * time.Second,
		IdleTimeout:  time.Duration(helpers.StrToInt(idleTimeoutSeconds)) * time.Second,
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
	port := s.util.Environment.GetEnv("server", "PORT")

	s.util.Log.Info("Server started", map[string]any{
		"address": s.server.Addr,
		"port":    port,
	})

	if err := s.server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *Server) WaitForShutdown() {
	var wait time.Duration
	gracefulTimeout := s.util.Environment.GetEnv("server", "GRACEFUL_TIMEOUT_SECONDS")
	GRACEFUL_TIMEOUT := time.Duration(helpers.StrToInt(gracefulTimeout)) * time.Second

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
