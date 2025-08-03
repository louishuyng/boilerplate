package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app     *fiber.App
	adapter *Adapter
}

func NewServer(adapter *Adapter) *Server {
	app := fiber.New()

	server := &Server{
		app:     app,
		adapter: adapter,
	}

	server.registerRoutes()

	return server
}

func (s *Server) registerRoutes() {
	s.app.Post("/payments", s.adapter.Charge)
	s.app.Get("/payments/:id", s.adapter.Get)
}

func (s *Server) Run(port int) error {
	if port == 0 {
		port = 8080
	}

	log.Printf("Server starting on port %d", port)
	return s.app.Listen(fmt.Sprintf(":%v", port))
}
