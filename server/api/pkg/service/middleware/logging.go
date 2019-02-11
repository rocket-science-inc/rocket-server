package middleware

import (
	grpc "rocket-server/server/api/pkg/transport/grpc/client"
)

type loggingMiddleware struct {
	logger log.Logger
	next   ApiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ApiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ApiService) ApiService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) AddEvent(ctx context.Context, e types.NewEvent) (event types.Event, err error) {
	defer func() {
		l.logger.Log("method", "AddEvent", "e", e, "event", event, "err", err)
	}()
	return l.next.AddEvent(ctx, e)
}

func (l loggingMiddleware) GetEvents(ctx context.Context) (events []types.Event, err error) {
	defer func() {
		l.logger.Log("method", "GetEvents", "events", events, "err", err)
	}()
	return l.next.GetEvents(ctx)
}
