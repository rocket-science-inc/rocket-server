package service

import (
	"context"

	types "rocket-server/server/api/pkg/types"
)

func (b *basicApiService) GetEvents(ctx context.Context) (events []types.Event, err error) {
	// TODO implement the business logic of GetEvents
	//panic("not implemented")
	
	var data []types.Event
	data = append(data,
		types.Event {
			ID: "1",
		},
	)
	return data, nil
}

func (l loggingMiddleware) GetEvents(ctx context.Context) (events []types.Event, err error) {
	defer func() {
		l.logger.Log("method", "GetEvents", "events", events, "err", err)
	}()
	return l.next.GetEvents(ctx)
}
