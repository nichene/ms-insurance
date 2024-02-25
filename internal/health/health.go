package health

import "context"

type healthCheckService struct{}

func NewHealthCheckService() Checks {
	return &healthCheckService{}
}

func (s *healthCheckService) Ping(ctx context.Context) error {
	return nil
}
