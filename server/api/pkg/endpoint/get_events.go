package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	service "rocket-server/server/api/pkg/service"
	types "rocket-server/server/api/pkg/types"
)

// GetEvents implements Service. Primarily useful in a client.
func (en Endpoints) GetEvents(ctx context.Context) (events []types.Event, err error) {
	request := GetEventsRequest{}
	response, err := en.GetEventsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetEventsResponse).Events, response.(GetEventsResponse).Err
}

// GetEventsRequest collects the request parameters for the GetEvents method.
type GetEventsRequest struct{}

// GetEventsResponse collects the response parameters for the GetEvents method.
type GetEventsResponse struct {
	Events []types.Event `json:"events"`
	Err    error         `json:"err"`
}

// MakeGetEventsEndpoint returns an endpoint that invokes GetEvents on the service.
func MakeGetEventsEndpoint(s service.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		events, err := s.GetEvents(ctx)
		return GetEventsResponse{
			Err:    err,
			Events: events,
		}, nil
	}
}

// Failed implements Failer.
func (r GetEventsResponse) Failed() error {
	return r.Err
}
