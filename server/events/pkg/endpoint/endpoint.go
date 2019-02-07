package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"

	service "rocket-server/server/events/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetEventsEndpoint endpoint.Endpoint
	AddEventEndpoint  endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.EventsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddEventEndpoint:  MakeAddEventEndpoint(s),
		GetEventsEndpoint: MakeGetEventsEndpoint(s),
	}
	for _, m := range mdw["GetEvents"] {
		eps.GetEventsEndpoint = m(eps.GetEventsEndpoint)
	}
	for _, m := range mdw["AddEvent"] {
		eps.AddEventEndpoint = m(eps.AddEventEndpoint)
	}
	return eps
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}
