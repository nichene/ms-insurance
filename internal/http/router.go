package http

import (
	"ms-insurance/internal/http/handlers"
	"ms-insurance/internal/http/routes"
)

func (s *Server) router() {
	routes.HealthRoute(s.httpServer, handlers.NewHealthHandler(s.service.Health))

	routes.ProductRoutes(s.httpServer, handlers.NewProductHandler(s.service.ProductService))

	routes.SwagRoute(s.httpServer)
}
