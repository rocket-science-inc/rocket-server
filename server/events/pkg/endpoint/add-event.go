package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	
	service "rocket-server/server/events/pkg/service"
	types "rocket-server/server/events/pkg/types"
)

// AddEventRequest collects the request parameters for the AddEvent method.
type AddEventRequest struct {
	E types.Event `json:"event"`
}

// AddEventResponse collects the response parameters for the AddEvent method.
type AddEventResponse struct {
	Event types.Event `json:"event"`
	Error error       `json:"error"`
}

// MakeAddEventEndpoint returns an endpoint that invokes AddEvent on the service.
func MakeAddEventEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddEventRequest)
		event, error := s.AddEvent(ctx, req.E)
		return AddEventResponse{
			Error: error,
			Event: event,
		}, nil
	}
}

// Failed implements Failer.
func (r AddEventResponse) Failed() error {
	return r.Error
}
