package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	types "rocket-server/server/users/pkg/types"
	service "rocket-server/server/users/pkg/service"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Users []types.User `json:"users"`
	Error error        `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, error := s.Get(ctx)
		return GetResponse{
			Error: error,
			Users: users,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	U types.User `json:"u"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	User  types.User `json:"user"`
	Error error      `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		user, error := s.Add(ctx, req.U)
		return AddResponse{
			Error: error,
			User:  user,
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
func (e Endpoints) Get(ctx context.Context) (users []types.User, error error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Users, response.(GetResponse).Error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, u types.User) (user types.User, error error) {
	request := AddRequest{U: u}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).User, response.(AddResponse).Error
}
