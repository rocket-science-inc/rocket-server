package service

import (
	"context"

	types "rocket-server/server/api/pkg/types"
)

// ApiService describes the service.
type ApiService interface {
	GetEvents(ctx context.Context) (events []types.Event, err error)
	AddEvent(ctx context.Context, e types.NewEvent) (event types.Event, err error)
}

type basicApiService struct{}

// NewBasicApiService returns a naive, stateless implementation of ApiService.
func NewBasicApiService() ApiService {
	return &basicApiService{}
}

// New returns a ApiService with all of the expected middleware wired in.
func New(middleware []Middleware) ApiService {
	var svc ApiService = NewBasicApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
