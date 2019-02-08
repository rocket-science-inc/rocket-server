package service

import (
	"context"

	types "rocket-server/server/api/pkg/types"
)

func (b *basicApiService) AddEvent(ctx context.Context, e types.Event) (event types.Event, err error) {
	// TODO implement the business logic of AddEvent
	panic("not implemented")
	return event, err
}

func (l loggingMiddleware) AddEvent(ctx context.Context, e types.Event) (event types.Event, err error) {
	defer func() {
		l.logger.Log("method", "AddEvent", "e", e, "event", event, "err", err)
	}()
	return l.next.AddEvent(ctx, e)
}
