package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	
	service "rocket-server/server/events/pkg/service"
	types "rocket-server/server/events/pkg/types"
)

// GetEventsRequest collects the request parameters for the GetEvents method.
type GetEventsRequest struct{}

// GetEventsResponse collects the response parameters for the GetEvents method.
type GetEventsResponse struct {
	Events []types.Event `json:"events"`
	Error  error         `json:"error"`
}

// MakeGetEventsEndpoint returns an endpoint that invokes GetEvents on the service.
func MakeGetEventsEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		events, error := s.GetEvents(ctx)
		return GetEventsResponse{
			Error:  error,
			Events: events,
		}, nil
	}
}

// Failed implements Failer.
func (r GetEventsResponse) Failed() error {
	return r.Error
}
