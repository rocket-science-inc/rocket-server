package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	service "rocket-server/server/api/pkg/service"
)

// GraphqlRequest collects the request parameters for the Graphql method.
type GraphqlRequest struct {

}

// GraphqlResponse collects the response parameters for the Graphql method.
type GraphqlResponse struct {

}

// MakeGraphqlEndpoint returns an endpoint that invokes Get on the service.
func MakeGraphqlEndpoint(s service.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return GraphqlResponse {

		}, nil
	}
}