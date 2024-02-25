package health

import "context"

type Checks interface {
	Ping(ctx context.Context) error
}
