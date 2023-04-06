package repo

import "context"

type Webhook interface {
	Get(ctx context.Context, eventID string) (bool, error)
	GetAndSet(ctx context.Context, eventID string) (bool, error)
	Delete(ctx context.Context, eventID string) error
}
