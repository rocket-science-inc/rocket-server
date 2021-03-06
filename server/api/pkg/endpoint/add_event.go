package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	service "rocket-server/server/api/pkg/service"
	types "rocket-server/server/api/pkg/types"
)

// AddEvent implements Service. Primarily useful in a client.
func (en Endpoints) AddEvent(ctx context.Context, e types.NewEvent) (event types.Event, err error) {
	request := AddEventRequest{E: e}
	response, err := en.AddEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddEventResponse).Event, response.(AddEventResponse).Err
}

// AddEventRequest collects the request parameters for the AddEvent method.
type AddEventRequest struct {
	E types.NewEvent `json:"e"`
}

// AddEventResponse collects the response parameters for the AddEvent method.
type AddEventResponse struct {
	Event types.Event `json:"event"`
	Err   error       `json:"err"`
}

// MakeAddEventEndpoint returns an endpoint that invokes AddEvent on the service.
func MakeAddEventEndpoint(s service.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddEventRequest)
		event, err := s.AddEvent(ctx, req.E)
		return AddEventResponse{
			Err:   err,
			Event: event,
		}, nil
	}
}

// Failed implements Failer.
func (r AddEventResponse) Failed() error {
	return r.Err
}
