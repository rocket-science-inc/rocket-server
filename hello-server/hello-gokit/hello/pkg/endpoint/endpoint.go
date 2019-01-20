package endpoint

import (
	"context"
	service "rocket-server/hello-server/hello-gokit/hello/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// SayHelloRequest collects the request parameters for the SayHello method.
type SayHelloRequest struct {
	S string `json:"s"`
}

// SayHelloResponse collects the response parameters for the SayHello method.
type SayHelloResponse struct {
	Str string `json:"str"`
	Err  error  `json:"err"`
}

// MakeSayHelloEndpoint returns an endpoint that invokes SayHello on the service.
func MakeSayHelloEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		str, err := s.SayHello(ctx, req.S)
		return SayHelloResponse {
			Err:  err,
			Str: str,
		}, nil
	}
}

// Failed implements Failer.
func (r SayHelloResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SayHello implements Service. Primarily useful in a client.
func (e Endpoints) SayHello(ctx context.Context, s string) (str string, err error) {
	request := SayHelloRequest{S: s}
	response, err := e.SayHelloEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SayHelloResponse).Str, response.(SayHelloResponse).Err
}
