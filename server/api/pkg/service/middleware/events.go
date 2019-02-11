package middleware

import (
	grpc "rocket-server/server/api/pkg/transport/grpc/client"
)

type eventsMiddleware struct {
	client 	grpc.EventsClient
	next	ApiService
}

// EventsMiddleware takes Events client as a dependency
// and returns a ApiService Middleware.
func EventsMiddleware(client grpc.EventsClient) Middleware {
	return func(service ApiService) ApiService {
		return &eventsServiceMiddleware{client, service}
	}
}

func (middleware eventsMiddleware) AddEvent(ctx context.Context, e types.NewEvent) (event types.Event, err error) {
	// TODO: add error check
	return middleware.client.service.AddEvent(ctx, e)
}

func (middleware eventsMiddleware) GetEvents(ctx context.Context) (events []types.Event, err error) {
	// TODO: add error check
	return middleware.client.service.GetEvents(ctx)
}

