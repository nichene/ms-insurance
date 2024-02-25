package main

import (
	"context"
	"log"
	"ms-insurance/config"
	"ms-insurance/internal/health"
	"ms-insurance/internal/http"
	"ms-insurance/internal/product"
	postgresRepo "ms-insurance/internal/product/postgres"
	postgres "ms-insurance/pkg"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadEnvVars()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := postgres.InitDatabase(ctx, cfg)

	httpService := http.NewService(
		health.NewHealthCheckService(),
		*product.NewService(
			postgresRepo.NewProductsRepository(db),
		),
	)

	srv, err := http.NewServer(cfg.Port, http.WithService(httpService))
	if err != nil {
		log.Fatal("API - Server down", err)
	}

	log.Default().Print("API - New server up on port ", cfg.Port)

	<-ctx.Done()

	stop()
	log.Default().Print("API - Shutting down gracefully")

	err = srv.Close()
	if err != nil {
		log.Fatal("API - Shutdown error", err)
	}

	log.Default().Print("API - exiting")
}
