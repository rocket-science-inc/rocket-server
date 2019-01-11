package service

import (
	"context"
	"rocket-server/server/events/pkg/types"
)

// EventsService describes the service.
type EventsService interface {
	Get(ctx context.Context) (events []types.Event, error error)
	Add(ctx context.Context, e types.Event) (event types.Event, error error)
}
