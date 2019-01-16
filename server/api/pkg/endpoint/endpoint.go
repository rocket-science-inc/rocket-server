package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	
	service "rocket-server/server/api/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GraphqlEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.ApiService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		GraphqlEndpoint: MakeGraphqlEndpoint(s),
	}
	for _, m := range mdw["Graphql"] {
		eps.GraphqlEndpoint = m(eps.GraphqlEndpoint)
	}
	return eps
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}
