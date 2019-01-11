package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "rocket-server/server/events/pkg/service"
	types "rocket-server/server/events/pkg/types"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Events []types.Event `json:"events"`
	Error  error         `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		events, error := s.Get(ctx)
		return GetResponse{
			Error:  error,
			Events: events,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	E types.Event `json:"e"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	Event types.Event `json:"event"`
	Error error       `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		event, error := s.Add(ctx, req.E)
		return AddResponse{
			Error: error,
			Event: event,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Error
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (en Endpoints) Get(ctx context.Context) (events []types.Event, error error) {
	request := GetRequest{}
	response, err := en.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Events, response.(GetResponse).Error
}

// Add implements Service. Primarily useful in a client.
func (en Endpoints) Add(ctx context.Context, e types.Event) (event types.Event, error error) {
	request := AddRequest{E: e}
	response, err := en.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).Event, response.(AddResponse).Error
}
