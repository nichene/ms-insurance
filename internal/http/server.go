package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	httpServer *fiber.App
	service    *Service
}

func NewServer(port string, options ...func(server *Server) error) (*Server, error) {
	server := &Server{}
	for _, option := range options {
		err := option(server)
		if err != nil {
			return nil, err
		}
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "ms-insurance",
	})

	server.httpServer = app
	server.router()

	go func() {
		port := fmt.Sprintf(":%s", port)
		err := server.httpServer.Listen(port)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("API - Server closed", err)
		}
	}()

	return server, nil
}

func WithService(service *Service) func(server *Server) error {
	return func(server *Server) error {
		server.service = service
		return nil
	}
}

func (s *Server) Close() error {
	return s.httpServer.Shutdown()
}
