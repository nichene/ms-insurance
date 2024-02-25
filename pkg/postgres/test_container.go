package postgres

import (
	"context"
	"fmt"
	"time"

	"ms-insurance/config"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestContainer struct {
	Instance testcontainers.Container
}

// NewPostgresTestContainer creates a postgres container using testcontainers-go
func NewPostgresTestContainer(ctx context.Context, cfg *config.Config) (*TestContainer, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"POSTGRES_USER":     cfg.DBUser,
			"POSTGRES_PASSWORD": cfg.DBPass,
			"POSTGRES_DB":       cfg.DBName,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		err = fmt.Errorf("failed to start postgres container: %w", err)
		return nil, err
	}

	return &TestContainer{
		Instance: postgres,
	}, nil
}

func (p *TestContainer) Port(ctx context.Context) (string, error) {
	port, err := p.Instance.MappedPort(ctx, "5432")
	if err != nil {
		err = fmt.Errorf("failed to get exposed postgres container port: %w", err)
		return "", err
	}

	return port.Port(), nil
}

func (p *TestContainer) Host(ctx context.Context) (string, error) {
	host, err := p.Instance.Host(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get postgres container ip: %w", err)
		return "", err
	}

	return host, nil
}
