package service

import (
	"context"

	db "rocket-server/server/events/pkg/database"
	types "rocket-server/server/events/pkg/types"
)

func (b *basicEventsService) GetEvents(ctx context.Context) (events []types.Event, err error) {
	// TODO implement the business logic of GetEvents

	dbConnectionString := "localhost:6379"
	client := db.NewClient(dbConnectionString)
	client.Ping()

	panic("not implemented")
	return events, err
}

func (l loggingMiddleware) GetEvents(ctx context.Context) (events []types.Event, err error) {
	defer func() {
		l.logger.Log("method", "GetEvents", "events", events, "err", err)
	}()
	return l.next.GetEvents(ctx)
}
