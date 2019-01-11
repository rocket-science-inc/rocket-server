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

// AddEventRequest collects the request parameters for the AddEvent method.
type AddEventRequest struct {
	E types.Event `json:"e"`
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

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetEvents implements Service. Primarily useful in a client.
func (en Endpoints) GetEvents(ctx context.Context) (events []types.Event, error error) {
	request := GetEventsRequest{}
	response, err := en.GetEventsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetEventsResponse).Events, response.(GetEventsResponse).Error
}

// AddEvent implements Service. Primarily useful in a client.
func (en Endpoints) AddEvent(ctx context.Context, e types.Event) (event types.Event, error error) {
	request := AddEventRequest{E: e}
	response, err := en.AddEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddEventResponse).Event, response.(AddEventResponse).Error
}
