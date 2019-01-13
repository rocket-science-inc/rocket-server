package service

import (
	"context"

	log "github.com/go-kit/kit/log"

	types "rocket-server/server/events/pkg/types"
	//db "rocket-server/server/events/pkg/db"
)

var logger log.Logger

// EventsService describes the service.
type EventsService interface {
	GetEvents(ctx context.Context) (events []types.Event, error error)
	AddEvent(ctx context.Context, e types.Event) (event types.Event, error error)
}

type basicEventsService struct{}

func (b *basicEventsService) GetEvents(ctx context.Context) (events []types.Event, error error) {
	// TODO implement the business logic of GetEvents
	//dbConnectionString := "localhost:6379"
	//logger.Log("Connect to redis server at address ", dbConnectionString)

	//client := db.NewClient(dbConnectionString)

	return events, error
}

func (b *basicEventsService) AddEvent(ctx context.Context, e types.Event) (event types.Event, error error) {
	// TODO implement the business logic of AddEvent
	return event, error
}

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
