package handler

import (
	"context"
	"errors"

	grpc "github.com/go-kit/kit/transport/grpc"
	
	endpoint "rocket-server/server/events/pkg/endpoint"
)

// makeGetEventsHandler creates the handler logic
func makeGetEventsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEventsEndpoint, decodeGetEventsRequest, encodeGetEventsResponse, options...)
}

// decodeGetEventsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeGetEventsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Decoder is not impelemented")
}

// encodeGetEventsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetEventsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Encoder is not impelemented")
}
