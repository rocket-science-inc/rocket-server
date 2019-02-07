package service

import (
	"context"

	types "rocket-server/server/events/pkg/types"
)

// EventsService describes the service.
type EventsService interface {
	GetEvents(ctx context.Context) (events []types.Event, err error)
	AddEvent(ctx context.Context, e types.Event) (event types.Event, err error)
}

type basicEventsService struct{}

// NewBasicEventsService returns a naive, stateless implementation of EventsService.
func NewBasicEventsService() EventsService {
	return &basicEventsService{}
}

// New returns a EventsService with all of the expected middleware wired in.
func New(middleware []Middleware) EventsService {
	var svc EventsService = NewBasicEventsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
